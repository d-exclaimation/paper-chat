//
//  Snackbar.swift
//  PaperChat
//
//  Created by Vincent on 11/17/21.
//

import SwiftUI

public struct Snackbar: View {
    
    public enum Variant {
        case info, warn, success, fatal
    }
    
    public var variant: Variant = .info
    public var message: String
    
    @Binding
    public var isPresented: Bool
    
    public var body: some View {
        VStack {
            Spacer()
            HStack {
                switch variant {
                case .info:
                    Image(systemName: "info.circle.fill")
                        .foregroundColor(.blue)
                case .warn:
                    Image(systemName: "exclamationmark.triangle.fill")
                        .foregroundColor(.yellow)
                case .success:
                    Image(systemName: "checkmark.circle.fill")
                        .foregroundColor(.green)
                case .fatal:
                    Image(systemName: "x.circle.fill")
                        .foregroundColor(.red)
                }
                Text(message)
                Spacer()
            }
            .padding()
            .background(background)
            .offset(x: 0, y: isPresented ? 0 : UIScreen.main.bounds.height)
        }
        .padding()
    }
    
    var background: some View {
        Rectangle()
            .colorInvert()
            .shadow(
                color: .black.opacity(0.3),
                radius: 4,
                x: 0,
                y: 2
            )
    }
}

struct Snackbar_Previews: PreviewProvider {
    static var previews: some View {
        Snackbar(message: "Snackbar", isPresented: Binding.constant(true))
    }
}
