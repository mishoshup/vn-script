import XCTest
import SwiftTreeSitter
import TreeSitterVnScript

final class TreeSitterVnScriptTests: XCTestCase {
    func testCanLoadGrammar() throws {
        let parser = Parser()
        let language = Language(language: tree_sitter_vn_script())
        XCTAssertNoThrow(try parser.setLanguage(language),
                         "Error loading VN Script grammar")
    }
}
