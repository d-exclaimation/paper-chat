//
//  InputBarView.swift
//  PaperChat
//
//  Created by Vincent on 11/17/21.
//

import SwiftUI

public struct InputBarView: View {
    @Binding
    public var input: String
    public var newMessage: () -> Void
    
    public var body: some View {
        HStack {
            quote
            textfield
            send
        }
        .frame(maxWidth: .infinity)
        .background(Color.paper)
    }
    
    var textfield: some View {
        TextField("Send a message", text: $input)
            .onSubmit(self.newMessage)
            .padding(.horizontal, 10)
            .padding(.vertical, .none)
    }
    
    var send: some View {
        Text("🕊")
            .font(.title)
            .padding(.trailing, 10)
            .button(action: self.newMessage)
    }
    
    var quote: some View {
        Image(systemName: "quote.opening")
            .font(.some(.title2))
            .opacity(0.25)
            .padding(.leading, 10)
    }
}

struct InputBarView_Previews: PreviewProvider {
    static var previews: some View {
        InputBarView(input: Binding.constant("")) {
            print("")
        }
    }
}
