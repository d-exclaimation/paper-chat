//
//  ChatMessageView.swift
//  PaperChat
//
//  Created by Vincent on 11/17/21.
//

import SwiftUI

public struct ChatMessageView<S>: View where S: ShapeStyle {
    
    public var content: String
    public var style: S
    public var onDelete: () -> Void
    
    @State
    private var isPopUp: Bool = false
    
    public var body: some View {
        VStack(alignment: .leading) {
            Text("Name")
                .font(.caption)
                .bold()
                .padding(.bottom, 0.25)
            if content.isEmpty {
                Text("empty message")
                    .foregroundColor(.primary.opacity(0.4))
                    .italic()
            } else {
                Text(content)
                    .textSelection(.enabled)
            }
        }
        .padding()
        .padding(.horizontal, 10)
        .background(style)
        .cornerRadius(10)
        .onLongPressGesture {
            isPopUp.toggle()
        }
        .confirmationDialog("Deleting this message?", isPresented: $isPopUp, titleVisibility: .visible) {
            Button("Delete", role: .destructive) {
                onDelete()
            }
        }
    }
}

struct ChatMessageView_Previews: PreviewProvider {
    static var previews: some View {
        ChatMessageView(content: "aa", style: Color.paperThin) {}
    }
}
