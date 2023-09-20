// Generated from c:\Users\gclka\antlr\grammar\tafexpr.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class TafexprLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, MUL=3, DIV=4, ADD=5, SUB=6, INTEGER=7, DOUBLE=8, WHITESPACE=9, 
		LBR=10, RBR=11, CON=12, DOLLAR=13, NUMBER=14, VARIABLE_NAME=15, PROP=16;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"T__0", "T__1", "MUL", "DIV", "ADD", "SUB", "INTEGER", "DOUBLE", "WHITESPACE", 
			"LBR", "RBR", "CON", "LOWERCASE", "UPPERCASE", "UNDERSCORE", "DOLLAR", 
			"SINGLELETTER", "VARIABLECON", "NUMBER", "VARIABLE_NAME", "PROP"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'('", "')'", "'*'", "'/'", "'+'", "'-'", null, null, null, "'['", 
			"']'", "'.'", "'$'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, "MUL", "DIV", "ADD", "SUB", "INTEGER", "DOUBLE", "WHITESPACE", 
			"LBR", "RBR", "CON", "DOLLAR", "NUMBER", "VARIABLE_NAME", "PROP"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}


	public TafexprLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "tafexpr.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\22\u0080\b\1\4\2"+
		"\t\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4"+
		"\13\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22"+
		"\t\22\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\3\2\3\2\3\3\3\3\3\4\3\4"+
		"\3\5\3\5\3\6\3\6\3\7\3\7\3\b\6\b;\n\b\r\b\16\b<\3\t\3\t\3\t\3\t\3\t\7"+
		"\tD\n\t\f\t\16\tG\13\t\3\t\3\t\6\tK\n\t\r\t\16\tL\5\tO\n\t\3\n\6\nR\n"+
		"\n\r\n\16\nS\3\n\3\n\3\13\3\13\3\f\3\f\3\r\3\r\3\16\3\16\3\17\3\17\3\20"+
		"\3\20\3\21\3\21\3\22\3\22\5\22h\n\22\3\23\3\23\3\24\3\24\3\25\3\25\3\25"+
		"\3\25\3\25\7\25s\n\25\f\25\16\25v\13\25\3\26\3\26\3\26\3\26\7\26|\n\26"+
		"\f\26\16\26\177\13\26\2\2\27\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13"+
		"\25\f\27\r\31\16\33\2\35\2\37\2!\17#\2%\2\'\20)\21+\22\3\2\b\3\2\62;\3"+
		"\2\62\62\3\2\63;\5\2\13\f\17\17\"\"\3\2c|\3\2C\\\2\u0086\2\3\3\2\2\2\2"+
		"\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2"+
		"\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2"+
		"!\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\3-\3\2\2\2\5/\3\2\2\2\7\61"+
		"\3\2\2\2\t\63\3\2\2\2\13\65\3\2\2\2\r\67\3\2\2\2\17:\3\2\2\2\21N\3\2\2"+
		"\2\23Q\3\2\2\2\25W\3\2\2\2\27Y\3\2\2\2\31[\3\2\2\2\33]\3\2\2\2\35_\3\2"+
		"\2\2\37a\3\2\2\2!c\3\2\2\2#g\3\2\2\2%i\3\2\2\2\'k\3\2\2\2)m\3\2\2\2+w"+
		"\3\2\2\2-.\7*\2\2.\4\3\2\2\2/\60\7+\2\2\60\6\3\2\2\2\61\62\7,\2\2\62\b"+
		"\3\2\2\2\63\64\7\61\2\2\64\n\3\2\2\2\65\66\7-\2\2\66\f\3\2\2\2\678\7/"+
		"\2\28\16\3\2\2\29;\t\2\2\2:9\3\2\2\2;<\3\2\2\2<:\3\2\2\2<=\3\2\2\2=\20"+
		"\3\2\2\2>?\t\3\2\2?@\7\60\2\2@O\t\3\2\2AE\t\4\2\2BD\t\2\2\2CB\3\2\2\2"+
		"DG\3\2\2\2EC\3\2\2\2EF\3\2\2\2FH\3\2\2\2GE\3\2\2\2HJ\7\60\2\2IK\t\2\2"+
		"\2JI\3\2\2\2KL\3\2\2\2LJ\3\2\2\2LM\3\2\2\2MO\3\2\2\2N>\3\2\2\2NA\3\2\2"+
		"\2O\22\3\2\2\2PR\t\5\2\2QP\3\2\2\2RS\3\2\2\2SQ\3\2\2\2ST\3\2\2\2TU\3\2"+
		"\2\2UV\b\n\2\2V\24\3\2\2\2WX\7]\2\2X\26\3\2\2\2YZ\7_\2\2Z\30\3\2\2\2["+
		"\\\7\60\2\2\\\32\3\2\2\2]^\t\6\2\2^\34\3\2\2\2_`\t\7\2\2`\36\3\2\2\2a"+
		"b\7a\2\2b \3\2\2\2cd\7&\2\2d\"\3\2\2\2eh\5\33\16\2fh\5\35\17\2ge\3\2\2"+
		"\2gf\3\2\2\2h$\3\2\2\2ij\7\60\2\2j&\3\2\2\2kl\5\17\b\2l(\3\2\2\2mn\5!"+
		"\21\2nt\5#\22\2os\5#\22\2ps\5\17\b\2qs\5\37\20\2ro\3\2\2\2rp\3\2\2\2r"+
		"q\3\2\2\2sv\3\2\2\2tr\3\2\2\2tu\3\2\2\2u*\3\2\2\2vt\3\2\2\2w}\5#\22\2"+
		"x|\5#\22\2y|\5\17\b\2z|\5\37\20\2{x\3\2\2\2{y\3\2\2\2{z\3\2\2\2|\177\3"+
		"\2\2\2}{\3\2\2\2}~\3\2\2\2~,\3\2\2\2\177}\3\2\2\2\r\2<ELNSgrt{}\3\b\2"+
		"\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}