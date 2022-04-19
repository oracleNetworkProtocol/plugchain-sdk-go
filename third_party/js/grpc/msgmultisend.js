// source: cosmos/bank/v1beta1/tx.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {missingRequire} reports error on implicit type usages.
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

goog.provide('proto.cosmos.bank.v1beta1.MsgMultiSend');

goog.require('jspb.BinaryReader');
goog.require('jspb.BinaryWriter');
goog.require('jspb.Message');
goog.require('proto.cosmos.bank.v1beta1.Input');
goog.require('proto.cosmos.bank.v1beta1.Output');

/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.cosmos.bank.v1beta1.MsgMultiSend = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.cosmos.bank.v1beta1.MsgMultiSend.repeatedFields_, null);
};
goog.inherits(proto.cosmos.bank.v1beta1.MsgMultiSend, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.cosmos.bank.v1beta1.MsgMultiSend.displayName = 'proto.cosmos.bank.v1beta1.MsgMultiSend';
}

/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.cosmos.bank.v1beta1.MsgMultiSend.repeatedFields_ = [1,2];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.cosmos.bank.v1beta1.MsgMultiSend.prototype.toObject = function(opt_includeInstance) {
  return proto.cosmos.bank.v1beta1.MsgMultiSend.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.cosmos.bank.v1beta1.MsgMultiSend} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.cosmos.bank.v1beta1.MsgMultiSend.toObject = function(includeInstance, msg) {
  var f, obj = {
    inputsList: jspb.Message.toObjectList(msg.getInputsList(),
    proto.cosmos.bank.v1beta1.Input.toObject, includeInstance),
    outputsList: jspb.Message.toObjectList(msg.getOutputsList(),
    proto.cosmos.bank.v1beta1.Output.toObject, includeInstance)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.cosmos.bank.v1beta1.MsgMultiSend}
 */
proto.cosmos.bank.v1beta1.MsgMultiSend.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.cosmos.bank.v1beta1.MsgMultiSend;
  return proto.cosmos.bank.v1beta1.MsgMultiSend.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.cosmos.bank.v1beta1.MsgMultiSend} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.cosmos.bank.v1beta1.MsgMultiSend}
 */
proto.cosmos.bank.v1beta1.MsgMultiSend.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.cosmos.bank.v1beta1.Input;
      reader.readMessage(value,proto.cosmos.bank.v1beta1.Input.deserializeBinaryFromReader);
      msg.addInputs(value);
      break;
    case 2:
      var value = new proto.cosmos.bank.v1beta1.Output;
      reader.readMessage(value,proto.cosmos.bank.v1beta1.Output.deserializeBinaryFromReader);
      msg.addOutputs(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.cosmos.bank.v1beta1.MsgMultiSend.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.cosmos.bank.v1beta1.MsgMultiSend.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.cosmos.bank.v1beta1.MsgMultiSend} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.cosmos.bank.v1beta1.MsgMultiSend.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getInputsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      proto.cosmos.bank.v1beta1.Input.serializeBinaryToWriter
    );
  }
  f = message.getOutputsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      2,
      f,
      proto.cosmos.bank.v1beta1.Output.serializeBinaryToWriter
    );
  }
};


/**
 * repeated Input inputs = 1;
 * @return {!Array<!proto.cosmos.bank.v1beta1.Input>}
 */
proto.cosmos.bank.v1beta1.MsgMultiSend.prototype.getInputsList = function() {
  return /** @type{!Array<!proto.cosmos.bank.v1beta1.Input>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.cosmos.bank.v1beta1.Input, 1));
};


/**
 * @param {!Array<!proto.cosmos.bank.v1beta1.Input>} value
 * @return {!proto.cosmos.bank.v1beta1.MsgMultiSend} returns this
*/
proto.cosmos.bank.v1beta1.MsgMultiSend.prototype.setInputsList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.cosmos.bank.v1beta1.Input=} opt_value
 * @param {number=} opt_index
 * @return {!proto.cosmos.bank.v1beta1.Input}
 */
proto.cosmos.bank.v1beta1.MsgMultiSend.prototype.addInputs = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.cosmos.bank.v1beta1.Input, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.cosmos.bank.v1beta1.MsgMultiSend} returns this
 */
proto.cosmos.bank.v1beta1.MsgMultiSend.prototype.clearInputsList = function() {
  return this.setInputsList([]);
};


/**
 * repeated Output outputs = 2;
 * @return {!Array<!proto.cosmos.bank.v1beta1.Output>}
 */
proto.cosmos.bank.v1beta1.MsgMultiSend.prototype.getOutputsList = function() {
  return /** @type{!Array<!proto.cosmos.bank.v1beta1.Output>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.cosmos.bank.v1beta1.Output, 2));
};


/**
 * @param {!Array<!proto.cosmos.bank.v1beta1.Output>} value
 * @return {!proto.cosmos.bank.v1beta1.MsgMultiSend} returns this
*/
proto.cosmos.bank.v1beta1.MsgMultiSend.prototype.setOutputsList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 2, value);
};


/**
 * @param {!proto.cosmos.bank.v1beta1.Output=} opt_value
 * @param {number=} opt_index
 * @return {!proto.cosmos.bank.v1beta1.Output}
 */
proto.cosmos.bank.v1beta1.MsgMultiSend.prototype.addOutputs = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 2, opt_value, proto.cosmos.bank.v1beta1.Output, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.cosmos.bank.v1beta1.MsgMultiSend} returns this
 */
proto.cosmos.bank.v1beta1.MsgMultiSend.prototype.clearOutputsList = function() {
  return this.setOutputsList([]);
};


