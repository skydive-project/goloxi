/*
 * Copyright (c) 2008 The Board of Trustees of The Leland Stanford Junior University
 * Copyright (c) 2011, 2012 Open Networking Foundation
 * Copyright 2013, Big Switch Networks, Inc. This library was generated by the LoxiGen Compiler.
 * Copyright 2018, Red Hat, Inc.
 */
// Automatically generated by LOXI from template module.go
// Do not modify

package of11

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/skydive-project/goloxi"
)

type Instruction struct {
	Type uint16
	Len  uint16
}

type IInstruction interface {
	goloxi.Serializable
	GetType() uint16
	GetLen() uint16
}

func (self *Instruction) GetType() uint16 {
	return self.Type
}

func (self *Instruction) SetType(v uint16) {
	self.Type = v
}

func (self *Instruction) GetLen() uint16 {
	return self.Len
}

func (self *Instruction) SetLen(v uint16) {
	self.Len = v
}

func (self *Instruction) Serialize(encoder *goloxi.Encoder) error {

	encoder.PutUint16(uint16(self.Type))
	encoder.PutUint16(uint16(self.Len))

	return nil
}

func DecodeInstruction(decoder *goloxi.Decoder) (IInstruction, error) {
	_instruction := &Instruction{}
	if decoder.Length() < 4 {
		return nil, fmt.Errorf("Instruction packet too short: %d < 4", decoder.Length())
	}
	_instruction.Type = uint16(decoder.ReadUint16())
	_instruction.Len = uint16(decoder.ReadUint16())
	oldDecoder := decoder
	defer func() { decoder = oldDecoder }()
	decoder = decoder.SliceDecoder(int(_instruction.Len), 2+2)

	switch _instruction.Type {
	case 1:
		return DecodeInstructionGotoTable(_instruction, decoder)
	case 2:
		return DecodeInstructionWriteMetadata(_instruction, decoder)
	case 3:
		return DecodeInstructionWriteActions(_instruction, decoder)
	case 4:
		return DecodeInstructionApplyActions(_instruction, decoder)
	case 5:
		return DecodeInstructionClearActions(_instruction, decoder)
	case 65535:
		return DecodeInstructionExperimenter(_instruction, decoder)
	default:
		return nil, fmt.Errorf("Invalid type '%d' for 'Instruction'", _instruction.Type)
	}
}

func NewInstruction(_type uint16) *Instruction {
	obj := &Instruction{}
	obj.Type = _type
	return obj
}

type InstructionApplyActions struct {
	*Instruction
	Actions []goloxi.IAction
}

type IInstructionApplyActions interface {
	IInstruction
	GetActions() []goloxi.IAction
}

func (self *InstructionApplyActions) GetActions() []goloxi.IAction {
	return self.Actions
}

func (self *InstructionApplyActions) SetActions(v []goloxi.IAction) {
	self.Actions = v
}

func (self *InstructionApplyActions) Serialize(encoder *goloxi.Encoder) error {
	if err := self.Instruction.Serialize(encoder); err != nil {
		return err
	}

	encoder.Write(bytes.Repeat([]byte{0}, 4))
	for _, obj := range self.Actions {
		if err := obj.Serialize(encoder); err != nil {
			return err
		}
	}

	binary.BigEndian.PutUint16(encoder.Bytes()[2:4], uint16(len(encoder.Bytes())))

	return nil
}

func DecodeInstructionApplyActions(parent *Instruction, decoder *goloxi.Decoder) (*InstructionApplyActions, error) {
	_instructionapplyactions := &InstructionApplyActions{Instruction: parent}
	decoder.Skip(4)

	for decoder.Length() >= 8 {
		item, err := DecodeAction(decoder)
		if err != nil {
			return nil, err
		}
		if item != nil {
			_instructionapplyactions.Actions = append(_instructionapplyactions.Actions, item)
		}
	}
	return _instructionapplyactions, nil
}

func NewInstructionApplyActions() *InstructionApplyActions {
	obj := &InstructionApplyActions{
		Instruction: NewInstruction(4),
	}
	return obj
}

type InstructionClearActions struct {
	*Instruction
}

type IInstructionClearActions interface {
	IInstruction
}

func (self *InstructionClearActions) Serialize(encoder *goloxi.Encoder) error {
	if err := self.Instruction.Serialize(encoder); err != nil {
		return err
	}

	binary.BigEndian.PutUint16(encoder.Bytes()[2:4], uint16(len(encoder.Bytes())))

	return nil
}

func DecodeInstructionClearActions(parent *Instruction, decoder *goloxi.Decoder) (*InstructionClearActions, error) {
	_instructionclearactions := &InstructionClearActions{Instruction: parent}
	if decoder.Length() < 4 {
		return nil, fmt.Errorf("InstructionClearActions packet too short: %d < 4", decoder.Length())
	}
	return _instructionclearactions, nil
}

func NewInstructionClearActions() *InstructionClearActions {
	obj := &InstructionClearActions{
		Instruction: NewInstruction(5),
	}
	return obj
}

type InstructionExperimenter struct {
	*Instruction
	Experimenter uint32
}

type IInstructionExperimenter interface {
	IInstruction
	GetExperimenter() uint32
}

func (self *InstructionExperimenter) GetExperimenter() uint32 {
	return self.Experimenter
}

func (self *InstructionExperimenter) SetExperimenter(v uint32) {
	self.Experimenter = v
}

func (self *InstructionExperimenter) Serialize(encoder *goloxi.Encoder) error {
	if err := self.Instruction.Serialize(encoder); err != nil {
		return err
	}

	encoder.PutUint32(uint32(self.Experimenter))

	return nil
}

func DecodeInstructionExperimenter(parent *Instruction, decoder *goloxi.Decoder) (IInstructionExperimenter, error) {
	_instructionexperimenter := &InstructionExperimenter{Instruction: parent}
	if decoder.Length() < 4 {
		return nil, fmt.Errorf("InstructionExperimenter packet too short: %d < 4", decoder.Length())
	}
	_instructionexperimenter.Experimenter = uint32(decoder.ReadUint32())
	return _instructionexperimenter, nil
}

func NewInstructionExperimenter(_experimenter uint32) *InstructionExperimenter {
	obj := &InstructionExperimenter{
		Instruction: NewInstruction(65535),
	}
	obj.Experimenter = _experimenter
	return obj
}

type InstructionGotoTable struct {
	*Instruction
	TableId uint8
}

type IInstructionGotoTable interface {
	IInstruction
	GetTableId() uint8
}

func (self *InstructionGotoTable) GetTableId() uint8 {
	return self.TableId
}

func (self *InstructionGotoTable) SetTableId(v uint8) {
	self.TableId = v
}

func (self *InstructionGotoTable) Serialize(encoder *goloxi.Encoder) error {
	if err := self.Instruction.Serialize(encoder); err != nil {
		return err
	}

	encoder.PutUint8(uint8(self.TableId))

	binary.BigEndian.PutUint16(encoder.Bytes()[2:4], uint16(len(encoder.Bytes())))

	return nil
}

func DecodeInstructionGotoTable(parent *Instruction, decoder *goloxi.Decoder) (*InstructionGotoTable, error) {
	_instructiongototable := &InstructionGotoTable{Instruction: parent}
	if decoder.Length() < 4 {
		return nil, fmt.Errorf("InstructionGotoTable packet too short: %d < 4", decoder.Length())
	}
	_instructiongototable.TableId = uint8(decoder.ReadByte())
	return _instructiongototable, nil
}

func NewInstructionGotoTable() *InstructionGotoTable {
	obj := &InstructionGotoTable{
		Instruction: NewInstruction(1),
	}
	return obj
}

type InstructionWriteActions struct {
	*Instruction
	Actions []goloxi.IAction
}

type IInstructionWriteActions interface {
	IInstruction
	GetActions() []goloxi.IAction
}

func (self *InstructionWriteActions) GetActions() []goloxi.IAction {
	return self.Actions
}

func (self *InstructionWriteActions) SetActions(v []goloxi.IAction) {
	self.Actions = v
}

func (self *InstructionWriteActions) Serialize(encoder *goloxi.Encoder) error {
	if err := self.Instruction.Serialize(encoder); err != nil {
		return err
	}

	encoder.Write(bytes.Repeat([]byte{0}, 4))
	for _, obj := range self.Actions {
		if err := obj.Serialize(encoder); err != nil {
			return err
		}
	}

	binary.BigEndian.PutUint16(encoder.Bytes()[2:4], uint16(len(encoder.Bytes())))

	return nil
}

func DecodeInstructionWriteActions(parent *Instruction, decoder *goloxi.Decoder) (*InstructionWriteActions, error) {
	_instructionwriteactions := &InstructionWriteActions{Instruction: parent}
	decoder.Skip(4)

	for decoder.Length() >= 8 {
		item, err := DecodeAction(decoder)
		if err != nil {
			return nil, err
		}
		if item != nil {
			_instructionwriteactions.Actions = append(_instructionwriteactions.Actions, item)
		}
	}
	return _instructionwriteactions, nil
}

func NewInstructionWriteActions() *InstructionWriteActions {
	obj := &InstructionWriteActions{
		Instruction: NewInstruction(3),
	}
	return obj
}

type InstructionWriteMetadata struct {
	*Instruction
	Metadata     uint64
	MetadataMask uint64
}

type IInstructionWriteMetadata interface {
	IInstruction
	GetMetadata() uint64
	GetMetadataMask() uint64
}

func (self *InstructionWriteMetadata) GetMetadata() uint64 {
	return self.Metadata
}

func (self *InstructionWriteMetadata) SetMetadata(v uint64) {
	self.Metadata = v
}

func (self *InstructionWriteMetadata) GetMetadataMask() uint64 {
	return self.MetadataMask
}

func (self *InstructionWriteMetadata) SetMetadataMask(v uint64) {
	self.MetadataMask = v
}

func (self *InstructionWriteMetadata) Serialize(encoder *goloxi.Encoder) error {
	if err := self.Instruction.Serialize(encoder); err != nil {
		return err
	}

	encoder.Write(bytes.Repeat([]byte{0}, 4))
	encoder.PutUint64(uint64(self.Metadata))
	encoder.PutUint64(uint64(self.MetadataMask))

	binary.BigEndian.PutUint16(encoder.Bytes()[2:4], uint16(len(encoder.Bytes())))

	return nil
}

func DecodeInstructionWriteMetadata(parent *Instruction, decoder *goloxi.Decoder) (*InstructionWriteMetadata, error) {
	_instructionwritemetadata := &InstructionWriteMetadata{Instruction: parent}
	if decoder.Length() < 16 {
		return nil, fmt.Errorf("InstructionWriteMetadata packet too short: %d < 16", decoder.Length())
	}
	decoder.Skip(4)
	_instructionwritemetadata.Metadata = uint64(decoder.ReadUint64())
	_instructionwritemetadata.MetadataMask = uint64(decoder.ReadUint64())
	return _instructionwritemetadata, nil
}

func NewInstructionWriteMetadata() *InstructionWriteMetadata {
	obj := &InstructionWriteMetadata{
		Instruction: NewInstruction(2),
	}
	return obj
}
