//
//  ContentView.swift
//  PaperChat
//
//  Created by Vincent on 11/16/21.
//

import Foundation
import SwiftUI

struct ContentView: View {
    @State
    private var rooms: [Room] = []
    
    var body: some View {
        NavigationView {
            List {
                ForEach(rooms) { room in
                    NavigationLink(destination: ChatScreen(room: room)) {
                        HStack {
                            Text("🛋 \(room.title)")
                            Spacer()
                            Text("\(room.member) 🕊")
                        }
                    }
                }
            }
            .navigationTitle("Paper Chat")
            .navigationBarItems(trailing: new)
        }
    }
    
    var new: some View {
        Image(systemName: "plus")
            .foregroundColor(.blue)
            .button(action: self.newRoom)
    }
    
    func newRoom() {
        withAnimation {
            let id = UUID()
            let room = Room(
                id: id,
                title: String(id.uuidString.suffix(6)),
                member: Int.random(in: 1...10)
            )
            rooms.append(room)
        }
    }
}

struct ContentView_Previews: PreviewProvider {
    static var previews: some View {
        ContentView()
    }
}
