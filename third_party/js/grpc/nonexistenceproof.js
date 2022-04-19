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

goog.provide('proto.ics23.NonExistenceProof');

goog.require('jspb.BinaryReader');
goog.require('jspb.BinaryWriter');
goog.require('jspb.Message');
goog.require('proto.ics23.ExistenceProof');

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
proto.ics23.NonExistenceProof = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.ics23.NonExistenceProof, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.ics23.NonExistenceProof.displayName = 'proto.ics23.NonExistenceProof';
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
proto.ics23.NonExistenceProof.prototype.toObject = function(opt_includeInstance) {
  return proto.ics23.NonExistenceProof.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.ics23.NonExistenceProof} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ics23.NonExistenceProof.toObject = function(includeInstance, msg) {
  var f, obj = {
    key: msg.getKey_asB64(),
    left: (f = msg.getLeft()) && proto.ics23.ExistenceProof.toObject(includeInstance, f),
    right: (f = msg.getRight()) && proto.ics23.ExistenceProof.toObject(includeInstance, f)
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
 * @return {!proto.ics23.NonExistenceProof}
 */
proto.ics23.NonExistenceProof.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.ics23.NonExistenceProof;
  return proto.ics23.NonExistenceProof.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.ics23.NonExistenceProof} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.ics23.NonExistenceProof}
 */
proto.ics23.NonExistenceProof.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!Uint8Array} */ (reader.readBytes());
      msg.setKey(value);
      break;
    case 2:
      var value = new proto.ics23.ExistenceProof;
      reader.readMessage(value,proto.ics23.ExistenceProof.deserializeBinaryFromReader);
      msg.setLeft(value);
      break;
    case 3:
      var value = new proto.ics23.ExistenceProof;
      reader.readMessage(value,proto.ics23.ExistenceProof.deserializeBinaryFromReader);
      msg.setRight(value);
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
proto.ics23.NonExistenceProof.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.ics23.NonExistenceProof.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.ics23.NonExistenceProof} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ics23.NonExistenceProof.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getKey_asU8();
  if (f.length > 0) {
    writer.writeBytes(
      1,
      f
    );
  }
  f = message.getLeft();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.ics23.ExistenceProof.serializeBinaryToWriter
    );
  }
  f = message.getRight();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      proto.ics23.ExistenceProof.serializeBinaryToWriter
    );
  }
};


/**
 * optional bytes key = 1;
 * @return {string}
 */
proto.ics23.NonExistenceProof.prototype.getKey = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * optional bytes key = 1;
 * This is a type-conversion wrapper around `getKey()`
 * @return {string}
 */
proto.ics23.NonExistenceProof.prototype.getKey_asB64 = function() {
  return /** @type {string} */ (jspb.Message.bytesAsB64(
      this.getKey()));
};


/**
 * optional bytes key = 1;
 * Note that Uint8Array is not supported on all browsers.
 * @see http://caniuse.com/Uint8Array
 * This is a type-conversion wrapper around `getKey()`
 * @return {!Uint8Array}
 */
proto.ics23.NonExistenceProof.prototype.getKey_asU8 = function() {
  return /** @type {!Uint8Array} */ (jspb.Message.bytesAsU8(
      this.getKey()));
};


/**
 * @param {!(string|Uint8Array)} value
 * @return {!proto.ics23.NonExistenceProof} returns this
 */
proto.ics23.NonExistenceProof.prototype.setKey = function(value) {
  return jspb.Message.setProto3BytesField(this, 1, value);
};


/**
 * optional ExistenceProof left = 2;
 * @return {?proto.ics23.ExistenceProof}
 */
proto.ics23.NonExistenceProof.prototype.getLeft = function() {
  return /** @type{?proto.ics23.ExistenceProof} */ (
    jspb.Message.getWrapperField(this, proto.ics23.ExistenceProof, 2));
};


/**
 * @param {?proto.ics23.ExistenceProof|undefined} value
 * @return {!proto.ics23.NonExistenceProof} returns this
*/
proto.ics23.NonExistenceProof.prototype.setLeft = function(value) {
  return jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.ics23.NonExistenceProof} returns this
 */
proto.ics23.NonExistenceProof.prototype.clearLeft = function() {
  return this.setLeft(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.ics23.NonExistenceProof.prototype.hasLeft = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional ExistenceProof right = 3;
 * @return {?proto.ics23.ExistenceProof}
 */
proto.ics23.NonExistenceProof.prototype.getRight = function() {
  return /** @type{?proto.ics23.ExistenceProof} */ (
    jspb.Message.getWrapperField(this, proto.ics23.ExistenceProof, 3));
};


/**
 * @param {?proto.ics23.ExistenceProof|undefined} value
 * @return {!proto.ics23.NonExistenceProof} returns this
*/
proto.ics23.NonExistenceProof.prototype.setRight = function(value) {
  return jspb.Message.setWrapperField(this, 3, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.ics23.NonExistenceProof} returns this
 */
proto.ics23.NonExistenceProof.prototype.clearRight = function() {
  return this.setRight(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.ics23.NonExistenceProof.prototype.hasRight = function() {
  return jspb.Message.getField(this, 3) != null;
};


