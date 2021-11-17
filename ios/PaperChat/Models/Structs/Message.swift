//
//  Message.swift
//  PaperChat
//
//  Created by Vincent on 11/16/21.
//

import Foundation

public struct Message: Identifiable {
    public let id: UUID = UUID()
    public var content: String
    public var sender: Sender = .me
}

public enum Sender {
    case me
    case other(name: String)
}
