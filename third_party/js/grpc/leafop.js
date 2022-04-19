// source: confio/proofs.proto
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

goog.provide('proto.ics23.LeafOp');

goog.require('jspb.BinaryReader');
goog.require('jspb.BinaryWriter');
goog.require('jspb.Message');

goog.forwardDeclare('proto.ics23.HashOp');
goog.forwardDeclare('proto.ics23.LengthOp');
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
proto.ics23.LeafOp = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.ics23.LeafOp, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.ics23.LeafOp.displayName = 'proto.ics23.LeafOp';
}



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
proto.ics23.LeafOp.prototype.toObject = function(opt_includeInstance) {
  return proto.ics23.LeafOp.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.ics23.LeafOp} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ics23.LeafOp.toObject = function(includeInstance, msg) {
  var f, obj = {
    hash: jspb.Message.getFieldWithDefault(msg, 1, 0),
    prehashKey: jspb.Message.getFieldWithDefault(msg, 2, 0),
    prehashValue: jspb.Message.getFieldWithDefault(msg, 3, 0),
    length: jspb.Message.getFieldWithDefault(msg, 4, 0),
    prefix: msg.getPrefix_asB64()
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
 * @return {!proto.ics23.LeafOp}
 */
proto.ics23.LeafOp.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.ics23.LeafOp;
  return proto.ics23.LeafOp.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.ics23.LeafOp} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.ics23.LeafOp}
 */
proto.ics23.LeafOp.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!proto.ics23.HashOp} */ (reader.readEnum());
      msg.setHash(value);
      break;
    case 2:
      var value = /** @type {!proto.ics23.HashOp} */ (reader.readEnum());
      msg.setPrehashKey(value);
      break;
    case 3:
      var value = /** @type {!proto.ics23.HashOp} */ (reader.readEnum());
      msg.setPrehashValue(value);
      break;
    case 4:
      var value = /** @type {!proto.ics23.LengthOp} */ (reader.readEnum());
      msg.setLength(value);
      break;
    case 5:
      var value = /** @type {!Uint8Array} */ (reader.readBytes());
      msg.setPrefix(value);
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
proto.ics23.LeafOp.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.ics23.LeafOp.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.ics23.LeafOp} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ics23.LeafOp.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getHash();
  if (f !== 0.0) {
    writer.writeEnum(
      1,
      f
    );
  }
  f = message.getPrehashKey();
  if (f !== 0.0) {
    writer.writeEnum(
      2,
      f
    );
  }
  f = message.getPrehashValue();
  if (f !== 0.0) {
    writer.writeEnum(
      3,
      f
    );
  }
  f = message.getLength();
  if (f !== 0.0) {
    writer.writeEnum(
      4,
      f
    );
  }
  f = message.getPrefix_asU8();
  if (f.length > 0) {
    writer.writeBytes(
      5,
      f
    );
  }
};


/**
 * optional HashOp hash = 1;
 * @return {!proto.ics23.HashOp}
 */
proto.ics23.LeafOp.prototype.getHash = function() {
  return /** @type {!proto.ics23.HashOp} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {!proto.ics23.HashOp} value
 * @return {!proto.ics23.LeafOp} returns this
 */
proto.ics23.LeafOp.prototype.setHash = function(value) {
  return jspb.Message.setProto3EnumField(this, 1, value);
};


/**
 * optional HashOp prehash_key = 2;
 * @return {!proto.ics23.HashOp}
 */
proto.ics23.LeafOp.prototype.getPrehashKey = function() {
  return /** @type {!proto.ics23.HashOp} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {!proto.ics23.HashOp} value
 * @return {!proto.ics23.LeafOp} returns this
 */
proto.ics23.LeafOp.prototype.setPrehashKey = function(value) {
  return jspb.Message.setProto3EnumField(this, 2, value);
};


/**
 * optional HashOp prehash_value = 3;
 * @return {!proto.ics23.HashOp}
 */
proto.ics23.LeafOp.prototype.getPrehashValue = function() {
  return /** @type {!proto.ics23.HashOp} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/**
 * @param {!proto.ics23.HashOp} value
 * @return {!proto.ics23.LeafOp} returns this
 */
proto.ics23.LeafOp.prototype.setPrehashValue = function(value) {
  return jspb.Message.setProto3EnumField(this, 3, value);
};


/**
 * optional LengthOp length = 4;
 * @return {!proto.ics23.LengthOp}
 */
proto.ics23.LeafOp.prototype.getLength = function() {
  return /** @type {!proto.ics23.LengthOp} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};


/**
 * @param {!proto.ics23.LengthOp} value
 * @return {!proto.ics23.LeafOp} returns this
 */
proto.ics23.LeafOp.prototype.setLength = function(value) {
  return jspb.Message.setProto3EnumField(this, 4, value);
};


/**
 * optional bytes prefix = 5;
 * @return {string}
 */
proto.ics23.LeafOp.prototype.getPrefix = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/**
 * optional bytes prefix = 5;
 * This is a type-conversion wrapper around `getPrefix()`
 * @return {string}
 */
proto.ics23.LeafOp.prototype.getPrefix_asB64 = function() {
  return /** @type {string} */ (jspb.Message.bytesAsB64(
      this.getPrefix()));
};


/**
 * optional bytes prefix = 5;
 * Note that Uint8Array is not supported on all browsers.
 * @see http://caniuse.com/Uint8Array
 * This is a type-conversion wrapper around `getPrefix()`
 * @return {!Uint8Array}
 */
proto.ics23.LeafOp.prototype.getPrefix_asU8 = function() {
  return /** @type {!Uint8Array} */ (jspb.Message.bytesAsU8(
      this.getPrefix()));
};


/**
 * @param {!(string|Uint8Array)} value
 * @return {!proto.ics23.LeafOp} returns this
 */
proto.ics23.LeafOp.prototype.setPrefix = function(value) {
  return jspb.Message.setProto3BytesField(this, 5, value);
};


