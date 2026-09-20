import Openreplay
import React

@objc(RnTrackerTouchViewManager)
class RnTrackerTouchViewManager: RCTViewManager {

  override func view() -> (RntrackerTouchView) {
    return RntrackerTouchView()
  }

  @objc override static func requiresMainQueueSetup() -> Bool {
    return true
  }
}

class RntrackerTouchView: UIView, UIGestureRecognizerDelegate {
  private var touchStart: CGPoint?

  override init(frame: CGRect) {
    super.init(frame: frame)
    setupRecognizers()
  }

  required init?(coder: NSCoder) {
    super.init(coder: coder)
    setupRecognizers()
  }

  private func setupRecognizers() {
    let tap = UITapGestureRecognizer(target: self, action: #selector(handleTap(_:)))
    tap.cancelsTouchesInView = false
    tap.delaysTouchesBegan = false
    tap.delaysTouchesEnded = false
    tap.delegate = self
    addGestureRecognizer(tap)

    let pan = UIPanGestureRecognizer(target: self, action: #selector(handlePan(_:)))
    pan.cancelsTouchesInView = false
    pan.delaysTouchesBegan = false
    pan.delaysTouchesEnded = false
    pan.delegate = self
    addGestureRecognizer(pan)
  }

  @objc private func handleTap(_ recognizer: UITapGestureRecognizer) {
    let point = recognizer.location(in: self)
    Analytics.shared.sendClick(label: "React-Native View", x: UInt64(point.x), y: UInt64(point.y))
  }

  @objc private func handlePan(_ recognizer: UIPanGestureRecognizer) {
    switch recognizer.state {
    case .began:
      touchStart = recognizer.location(in: self)
    case .ended:
      guard let startPoint = touchStart else { return }
      let endPoint = recognizer.location(in: self)
      touchStart = nil

      let deltaX = endPoint.x - startPoint.x
      let deltaY = endPoint.y - startPoint.y
      let distance = sqrt(deltaX * deltaX + deltaY * deltaY)

      if distance > 10 {
        let direction = abs(deltaX) > abs(deltaY) ? (deltaX > 0 ? "right" : "left") : (deltaY > 0 ? "down" : "up")
        Analytics.shared.sendSwipe(label: "React-Native View", x: UInt64(endPoint.x), y: UInt64(endPoint.y), direction: direction)
      }
    case .cancelled, .failed:
      touchStart = nil
    default:
      break
    }
  }

  func gestureRecognizer(_ gestureRecognizer: UIGestureRecognizer,
                         shouldRecognizeSimultaneouslyWith otherGestureRecognizer: UIGestureRecognizer) -> Bool {
    return true
  }
}
