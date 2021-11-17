//
//  ChatScreen.swift
//  PaperChat
//
//  Created by Vincent on 11/17/21.
//

import SwiftUI

public struct ChatScreen: View {
    
    public var room: Room
    
    @State
    private var chat = [Message]()
    @State
    private var input = ""
    @State
    private var isShowingInvite: Bool = false
    @State
    private var isSnackbar: Bool = false
    
    public var body: some View {
        VStack {
            ScrollView(.vertical, showsIndicators: false) {
                bar
                ForEach(chat.reversed()) { message in
                    HStack {
                        if case .other(name: _) = message.sender {
                            Spacer()
                            ChatMessageView(content: message.content, style: .bar) {
                                guard let index = chat.firstIndex(where: { $0.id == message.id }) else { return }
                                withAnimation {
                                    let _ = chat.remove(at: index)
                                }
                            }
                        } else {
                            ChatMessageView(content: message.content, style: Color.paperThin) {
                                guard let index = chat.firstIndex(where:  { $0.id == message.id }) else { return }
                                withAnimation {
                                    let _ = chat.remove(at: index)
                                }
                            }
                            Spacer()
                        }
                        
                    }
                    .padding(.horizontal, 5)
                    .flippedUpsideDown()
                }
            }
            .flippedUpsideDown()
            
            InputBarView(input: $input, newMessage: self.newMessage)
        }
        .navigationBarTitle(Text(room.title))
        .navigationBarTitleDisplayMode(.inline)
        .navigationBarItems(trailing: invite)
        .popover(isPresented: $isShowingInvite) {
            InviteCodeView(inviteId: room.id)
        }
    }

    
    var bar: some View {
        Rectangle()
            .size(width: .infinity, height: 0)
    }
    
    var invite: some View {
        Button("\(room.member)", role: .cancel) {
            isShowingInvite = true
        }
    }
    
    
    private func newMessage() {
        withAnimation {
            chat.append(Message(
                content: input,
                sender: Int.random(in: 0...1) == 0 ? .me : .other(name: UUID().uuidString)
            ))
            input.reset()
        }
    }
}

struct ChatScreen_Previews: PreviewProvider {
    static var previews: some View {
        ChatScreen(room: Room(title: "Hello", member: 0))
    }
}
