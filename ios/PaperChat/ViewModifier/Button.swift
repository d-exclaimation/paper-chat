//
//  Button.swift
//  PaperChat
//
//  Created by Vincent on 11/16/21.
//

import SwiftUI

extension View {
    public func button(action: @escaping () -> Void) -> some View {
        Button(action: action) { self }
    }
}
