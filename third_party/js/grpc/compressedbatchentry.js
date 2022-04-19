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

goog.provide('proto.ics23.CompressedBatchEntry');
goog.provide('proto.ics23.CompressedBatchEntry.ProofCase');

goog.require('jspb.BinaryReader');
goog.require('jspb.BinaryWriter');
goog.require('jspb.Message');
goog.require('proto.ics23.CompressedExistenceProof');
goog.require('proto.ics23.CompressedNonExistenceProof');

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
proto.ics23.CompressedBatchEntry = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, proto.ics23.CompressedBatchEntry.oneofGroups_);
};
goog.inherits(proto.ics23.CompressedBatchEntry, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.ics23.CompressedBatchEntry.displayName = 'proto.ics23.CompressedBatchEntry';
}

/**
 * Oneof group definitions for this message. Each group defines the field
 * numbers belonging to that group. When of these fields' value is set, all
 * other fields in the group are cleared. During deserialization, if multiple
 * fields are encountered for a group, only the last value seen will be kept.
 * @private {!Array<!Array<number>>}
 * @const
 */
proto.ics23.CompressedBatchEntry.oneofGroups_ = [[1,2]];

/**
 * @enum {number}
 */
proto.ics23.CompressedBatchEntry.ProofCase = {
  PROOF_NOT_SET: 0,
  EXIST: 1,
  NONEXIST: 2
};

/**
 * @return {proto.ics23.CompressedBatchEntry.ProofCase}
 */
proto.ics23.CompressedBatchEntry.prototype.getProofCase = function() {
  return /** @type {proto.ics23.CompressedBatchEntry.ProofCase} */(jspb.Message.computeOneofCase(this, proto.ics23.CompressedBatchEntry.oneofGroups_[0]));
};



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
proto.ics23.CompressedBatchEntry.prototype.toObject = function(opt_includeInstance) {
  return proto.ics23.CompressedBatchEntry.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.ics23.CompressedBatchEntry} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ics23.CompressedBatchEntry.toObject = function(includeInstance, msg) {
  var f, obj = {
    exist: (f = msg.getExist()) && proto.ics23.CompressedExistenceProof.toObject(includeInstance, f),
    nonexist: (f = msg.getNonexist()) && proto.ics23.CompressedNonExistenceProof.toObject(includeInstance, f)
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
 * @return {!proto.ics23.CompressedBatchEntry}
 */
proto.ics23.CompressedBatchEntry.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.ics23.CompressedBatchEntry;
  return proto.ics23.CompressedBatchEntry.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.ics23.CompressedBatchEntry} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.ics23.CompressedBatchEntry}
 */
proto.ics23.CompressedBatchEntry.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.ics23.CompressedExistenceProof;
      reader.readMessage(value,proto.ics23.CompressedExistenceProof.deserializeBinaryFromReader);
      msg.setExist(value);
      break;
    case 2:
      var value = new proto.ics23.CompressedNonExistenceProof;
      reader.readMessage(value,proto.ics23.CompressedNonExistenceProof.deserializeBinaryFromReader);
      msg.setNonexist(value);
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
proto.ics23.CompressedBatchEntry.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.ics23.CompressedBatchEntry.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.ics23.CompressedBatchEntry} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ics23.CompressedBatchEntry.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getExist();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.ics23.CompressedExistenceProof.serializeBinaryToWriter
    );
  }
  f = message.getNonexist();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.ics23.CompressedNonExistenceProof.serializeBinaryToWriter
    );
  }
};


/**
 * optional CompressedExistenceProof exist = 1;
 * @return {?proto.ics23.CompressedExistenceProof}
 */
proto.ics23.CompressedBatchEntry.prototype.getExist = function() {
  return /** @type{?proto.ics23.CompressedExistenceProof} */ (
    jspb.Message.getWrapperField(this, proto.ics23.CompressedExistenceProof, 1));
};


/**
 * @param {?proto.ics23.CompressedExistenceProof|undefined} value
 * @return {!proto.ics23.CompressedBatchEntry} returns this
*/
proto.ics23.CompressedBatchEntry.prototype.setExist = function(value) {
  return jspb.Message.setOneofWrapperField(this, 1, proto.ics23.CompressedBatchEntry.oneofGroups_[0], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.ics23.CompressedBatchEntry} returns this
 */
proto.ics23.CompressedBatchEntry.prototype.clearExist = function() {
  return this.setExist(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.ics23.CompressedBatchEntry.prototype.hasExist = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional CompressedNonExistenceProof nonexist = 2;
 * @return {?proto.ics23.CompressedNonExistenceProof}
 */
proto.ics23.CompressedBatchEntry.prototype.getNonexist = function() {
  return /** @type{?proto.ics23.CompressedNonExistenceProof} */ (
    jspb.Message.getWrapperField(this, proto.ics23.CompressedNonExistenceProof, 2));
};


/**
 * @param {?proto.ics23.CompressedNonExistenceProof|undefined} value
 * @return {!proto.ics23.CompressedBatchEntry} returns this
*/
proto.ics23.CompressedBatchEntry.prototype.setNonexist = function(value) {
  return jspb.Message.setOneofWrapperField(this, 2, proto.ics23.CompressedBatchEntry.oneofGroups_[0], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.ics23.CompressedBatchEntry} returns this
 */
proto.ics23.CompressedBatchEntry.prototype.clearNonexist = function() {
  return this.setNonexist(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.ics23.CompressedBatchEntry.prototype.hasNonexist = function() {
  return jspb.Message.getField(this, 2) != null;
};


