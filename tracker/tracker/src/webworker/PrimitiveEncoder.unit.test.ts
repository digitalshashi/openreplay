import { describe, expect, test, jest, beforeEach, afterEach } from '@jest/globals'
import PrimitiveEncoder from './PrimitiveEncoder.js'

describe('PrimitiveEncoder', () => {
  test('initial state', () => {
    const enc = new PrimitiveEncoder(10)

    expect(enc.getCurrentOffset()).toBe(0)
    expect(enc.isEmpty).toBe(true)
    expect(enc.flush().length).toBe(0)
  })

  test('skip()', () => {
    const enc = new PrimitiveEncoder(10)
    enc.skip(5)
    expect(enc.getCurrentOffset()).toBe(5)
    expect(enc.isEmpty).toBe(false)
  })

  test('checkpoint()', () => {
    const enc = new PrimitiveEncoder(10)
    enc.skip(5)
    enc.checkpoint()
    expect(enc.flush().length).toBe(5)
  })

  test('boolean(true)', () => {
    const enc = new PrimitiveEncoder(10)
    enc.boolean(true)
    enc.checkpoint()
    const bytes = enc.flush()
    expect(bytes.length).toBe(1)
    expect(bytes[0]).toBe(1)
  })
  test('boolean(false)', () => {
    const enc = new PrimitiveEncoder(10)
    enc.boolean(false)
    enc.checkpoint()
    const bytes = enc.flush()
    expect(bytes.length).toBe(1)
    expect(bytes[0]).toBe(0)
  })
  // TODO: test correct enc/dec on a top level(?) with player(PrimitiveReader.ts)/tracker(PrimitiveEncoder.ts)

  test('buffer oveflow with string()', () => {
    const N = 10
    const enc = new PrimitiveEncoder(N)
    const wasWritten = enc.string('long string'.repeat(N))
    expect(wasWritten).toBe(false)
  })
  test('buffer oveflow with uint()', () => {
    const enc = new PrimitiveEncoder(1)
    const wasWritten = enc.uint(Number.MAX_SAFE_INTEGER)
    expect(wasWritten).toBe(false)
  })
  test('buffer oveflow with boolean()', () => {
    const enc = new PrimitiveEncoder(1)
    let wasWritten = enc.boolean(true)
    expect(wasWritten).toBe(true)
    wasWritten = enc.boolean(true)
    expect(wasWritten).toBe(false)
  })

  describe('rewind()', () => {
    test('rolls back offset and checkpoint to a saved state', () => {
      const e = new PrimitiveEncoder(64)
      e.uint(1)
      e.uint(2)
      e.checkpoint()
      const savedOffset = e.getCurrentOffset()
      const savedCheckpoint = e.getCurrentCheckpoint()

      e.uint(3)
      e.uint(4)
      e.checkpoint()
      expect(e.getCurrentOffset()).toBeGreaterThan(savedOffset)

      e.rewind(savedOffset, savedCheckpoint)
      expect(e.getCurrentOffset()).toBe(savedOffset)
      expect(e.getCurrentCheckpoint()).toBe(savedCheckpoint)

      const out = e.flush()
      expect(Array.from(out)).toEqual([1, 2])
    })

    test('refuses to advance forward (no-op on bad input)', () => {
      const e = new PrimitiveEncoder(64)
      e.uint(7)
      const offset = e.getCurrentOffset()
      e.rewind(offset + 5, 100)
      expect(e.getCurrentOffset()).toBe(offset)
      expect(e.getCurrentCheckpoint()).toBe(0)
    })

    test('rewind to (0,0) is equivalent to reset', () => {
      const e = new PrimitiveEncoder(64)
      e.uint(1)
      e.checkpoint()
      e.uint(2)
      e.rewind(0, 0)
      expect(e.getCurrentOffset()).toBe(0)
      expect(e.getCurrentCheckpoint()).toBe(0)
    })
  })
})
