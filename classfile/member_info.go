package classfile

type MenberInfo struct {
	cp               ConstantPool
	accessFlags      uint16
	nameIndex        uint16
	descriptionIndex uint16
	attributes       []AttributeInfo
}

func readMembers(reader *ClassReader, cp ConstantPool) []*MenberInfo {
	memberCount := reader.readUint16()
	members := make([]*MenberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}

func readMember(reader *ClassReader, cp ConstantPool) *MenberInfo {
	return &MenberInfo{
		cp: cp
		accessFlags:reader.readUint16()
		nameIndex:reader.readUint16()
		descriptionIndex:reader.readUint16()
		attributes:readAttributes(reader,cp)
	}
}
func (self *MenberInfo) AccessFlags() uint16 {
	return self.AccessFlags
}

func (self *MenberInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}

func (self *MenberInfo) Descriptor() string {
	return self.cp.getUtf8(self.descriptionIndex)
}
