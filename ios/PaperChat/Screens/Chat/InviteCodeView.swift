//
//  InviteCodeView.swift
//  PaperChat
//
//  Created by Vincent on 11/17/21.
//

import SwiftUI

public struct InviteCodeView: View {
    public var inviteId: UUID
    
    @State
    private var isSnackbar: Bool = false
    
    public var body: some View {
        ZStack {
            VStack {
                Text("Room invite code:")
                    .font(.title)
                    .bold()
                    .padding(.bottom, .none)
                Text("\(inviteId)")
                    .textSelection(.enabled)
                    .opacity(0.6)
                    .onTapGesture {
                        UIPasteboard.general.setValue(inviteId.uuidString, forPasteboardType: "public.plain-text")
                        withAnimation {
                            isSnackbar = true
                        }
                    }
            }
            .padding()
            .padding(.horizontal, 12)
            
            Snackbar(message: "Copied to Clipboard", isPresented: $isSnackbar)
        }
    }
}

struct InviteCodeView_Previews: PreviewProvider {
    static var previews: some View {
        InviteCodeView(inviteId: UUID())
    }
}
