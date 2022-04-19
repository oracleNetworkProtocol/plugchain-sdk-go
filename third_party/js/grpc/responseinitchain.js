// source: tendermint/abci/types.proto
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

goog.provide('proto.tendermint.abci.ResponseInitChain');

goog.require('jspb.BinaryReader');
goog.require('jspb.BinaryWriter');
goog.require('jspb.Message');
goog.require('proto.tendermint.abci.ConsensusParams');
goog.require('proto.tendermint.abci.ValidatorUpdate');

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
proto.tendermint.abci.ResponseInitChain = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.tendermint.abci.ResponseInitChain.repeatedFields_, null);
};
goog.inherits(proto.tendermint.abci.ResponseInitChain, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.tendermint.abci.ResponseInitChain.displayName = 'proto.tendermint.abci.ResponseInitChain';
}

/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.tendermint.abci.ResponseInitChain.repeatedFields_ = [2];



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
proto.tendermint.abci.ResponseInitChain.prototype.toObject = function(opt_includeInstance) {
  return proto.tendermint.abci.ResponseInitChain.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.tendermint.abci.ResponseInitChain} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.tendermint.abci.ResponseInitChain.toObject = function(includeInstance, msg) {
  var f, obj = {
    consensusParams: (f = msg.getConsensusParams()) && proto.tendermint.abci.ConsensusParams.toObject(includeInstance, f),
    validatorsList: jspb.Message.toObjectList(msg.getValidatorsList(),
    proto.tendermint.abci.ValidatorUpdate.toObject, includeInstance),
    appHash: msg.getAppHash_asB64()
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
 * @return {!proto.tendermint.abci.ResponseInitChain}
 */
proto.tendermint.abci.ResponseInitChain.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.tendermint.abci.ResponseInitChain;
  return proto.tendermint.abci.ResponseInitChain.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.tendermint.abci.ResponseInitChain} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.tendermint.abci.ResponseInitChain}
 */
proto.tendermint.abci.ResponseInitChain.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.tendermint.abci.ConsensusParams;
      reader.readMessage(value,proto.tendermint.abci.ConsensusParams.deserializeBinaryFromReader);
      msg.setConsensusParams(value);
      break;
    case 2:
      var value = new proto.tendermint.abci.ValidatorUpdate;
      reader.readMessage(value,proto.tendermint.abci.ValidatorUpdate.deserializeBinaryFromReader);
      msg.addValidators(value);
      break;
    case 3:
      var value = /** @type {!Uint8Array} */ (reader.readBytes());
      msg.setAppHash(value);
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
proto.tendermint.abci.ResponseInitChain.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.tendermint.abci.ResponseInitChain.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.tendermint.abci.ResponseInitChain} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.tendermint.abci.ResponseInitChain.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getConsensusParams();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.tendermint.abci.ConsensusParams.serializeBinaryToWriter
    );
  }
  f = message.getValidatorsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      2,
      f,
      proto.tendermint.abci.ValidatorUpdate.serializeBinaryToWriter
    );
  }
  f = message.getAppHash_asU8();
  if (f.length > 0) {
    writer.writeBytes(
      3,
      f
    );
  }
};


/**
 * optional ConsensusParams consensus_params = 1;
 * @return {?proto.tendermint.abci.ConsensusParams}
 */
proto.tendermint.abci.ResponseInitChain.prototype.getConsensusParams = function() {
  return /** @type{?proto.tendermint.abci.ConsensusParams} */ (
    jspb.Message.getWrapperField(this, proto.tendermint.abci.ConsensusParams, 1));
};


/**
 * @param {?proto.tendermint.abci.ConsensusParams|undefined} value
 * @return {!proto.tendermint.abci.ResponseInitChain} returns this
*/
proto.tendermint.abci.ResponseInitChain.prototype.setConsensusParams = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.tendermint.abci.ResponseInitChain} returns this
 */
proto.tendermint.abci.ResponseInitChain.prototype.clearConsensusParams = function() {
  return this.setConsensusParams(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.tendermint.abci.ResponseInitChain.prototype.hasConsensusParams = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * repeated ValidatorUpdate validators = 2;
 * @return {!Array<!proto.tendermint.abci.ValidatorUpdate>}
 */
proto.tendermint.abci.ResponseInitChain.prototype.getValidatorsList = function() {
  return /** @type{!Array<!proto.tendermint.abci.ValidatorUpdate>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.tendermint.abci.ValidatorUpdate, 2));
};


/**
 * @param {!Array<!proto.tendermint.abci.ValidatorUpdate>} value
 * @return {!proto.tendermint.abci.ResponseInitChain} returns this
*/
proto.tendermint.abci.ResponseInitChain.prototype.setValidatorsList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 2, value);
};


/**
 * @param {!proto.tendermint.abci.ValidatorUpdate=} opt_value
 * @param {number=} opt_index
 * @return {!proto.tendermint.abci.ValidatorUpdate}
 */
proto.tendermint.abci.ResponseInitChain.prototype.addValidators = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 2, opt_value, proto.tendermint.abci.ValidatorUpdate, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.tendermint.abci.ResponseInitChain} returns this
 */
proto.tendermint.abci.ResponseInitChain.prototype.clearValidatorsList = function() {
  return this.setValidatorsList([]);
};


/**
 * optional bytes app_hash = 3;
 * @return {string}
 */
proto.tendermint.abci.ResponseInitChain.prototype.getAppHash = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * optional bytes app_hash = 3;
 * This is a type-conversion wrapper around `getAppHash()`
 * @return {string}
 */
proto.tendermint.abci.ResponseInitChain.prototype.getAppHash_asB64 = function() {
  return /** @type {string} */ (jspb.Message.bytesAsB64(
      this.getAppHash()));
};


/**
 * optional bytes app_hash = 3;
 * Note that Uint8Array is not supported on all browsers.
 * @see http://caniuse.com/Uint8Array
 * This is a type-conversion wrapper around `getAppHash()`
 * @return {!Uint8Array}
 */
proto.tendermint.abci.ResponseInitChain.prototype.getAppHash_asU8 = function() {
  return /** @type {!Uint8Array} */ (jspb.Message.bytesAsU8(
      this.getAppHash()));
};


/**
 * @param {!(string|Uint8Array)} value
 * @return {!proto.tendermint.abci.ResponseInitChain} returns this
 */
proto.tendermint.abci.ResponseInitChain.prototype.setAppHash = function(value) {
  return jspb.Message.setProto3BytesField(this, 3, value);
};


