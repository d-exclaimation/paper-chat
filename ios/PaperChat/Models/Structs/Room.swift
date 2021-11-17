//
//  Room.swift
//  PaperChat
//
//  Created by Vincent on 11/17/21.
//

import Foundation

public struct Room: Identifiable {
    public var id: UUID = UUID()
    public var title: String
    public var member: Int
}
