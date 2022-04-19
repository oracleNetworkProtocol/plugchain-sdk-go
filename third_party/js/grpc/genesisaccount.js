// source: ethermint/evm/v1/genesis.proto
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

goog.provide('proto.ethermint.evm.v1.GenesisAccount');

goog.require('jspb.BinaryReader');
goog.require('jspb.BinaryWriter');
goog.require('jspb.Message');
goog.require('proto.ethermint.evm.v1.State');

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
proto.ethermint.evm.v1.GenesisAccount = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.ethermint.evm.v1.GenesisAccount.repeatedFields_, null);
};
goog.inherits(proto.ethermint.evm.v1.GenesisAccount, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.ethermint.evm.v1.GenesisAccount.displayName = 'proto.ethermint.evm.v1.GenesisAccount';
}

/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.ethermint.evm.v1.GenesisAccount.repeatedFields_ = [3];



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
proto.ethermint.evm.v1.GenesisAccount.prototype.toObject = function(opt_includeInstance) {
  return proto.ethermint.evm.v1.GenesisAccount.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.ethermint.evm.v1.GenesisAccount} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ethermint.evm.v1.GenesisAccount.toObject = function(includeInstance, msg) {
  var f, obj = {
    address: jspb.Message.getFieldWithDefault(msg, 1, ""),
    code: jspb.Message.getFieldWithDefault(msg, 2, ""),
    storageList: jspb.Message.toObjectList(msg.getStorageList(),
    proto.ethermint.evm.v1.State.toObject, includeInstance)
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
 * @return {!proto.ethermint.evm.v1.GenesisAccount}
 */
proto.ethermint.evm.v1.GenesisAccount.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.ethermint.evm.v1.GenesisAccount;
  return proto.ethermint.evm.v1.GenesisAccount.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.ethermint.evm.v1.GenesisAccount} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.ethermint.evm.v1.GenesisAccount}
 */
proto.ethermint.evm.v1.GenesisAccount.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setAddress(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setCode(value);
      break;
    case 3:
      var value = new proto.ethermint.evm.v1.State;
      reader.readMessage(value,proto.ethermint.evm.v1.State.deserializeBinaryFromReader);
      msg.addStorage(value);
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
proto.ethermint.evm.v1.GenesisAccount.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.ethermint.evm.v1.GenesisAccount.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.ethermint.evm.v1.GenesisAccount} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ethermint.evm.v1.GenesisAccount.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getAddress();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getCode();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getStorageList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      3,
      f,
      proto.ethermint.evm.v1.State.serializeBinaryToWriter
    );
  }
};


/**
 * optional string address = 1;
 * @return {string}
 */
proto.ethermint.evm.v1.GenesisAccount.prototype.getAddress = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.ethermint.evm.v1.GenesisAccount} returns this
 */
proto.ethermint.evm.v1.GenesisAccount.prototype.setAddress = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string code = 2;
 * @return {string}
 */
proto.ethermint.evm.v1.GenesisAccount.prototype.getCode = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.ethermint.evm.v1.GenesisAccount} returns this
 */
proto.ethermint.evm.v1.GenesisAccount.prototype.setCode = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * repeated State storage = 3;
 * @return {!Array<!proto.ethermint.evm.v1.State>}
 */
proto.ethermint.evm.v1.GenesisAccount.prototype.getStorageList = function() {
  return /** @type{!Array<!proto.ethermint.evm.v1.State>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.ethermint.evm.v1.State, 3));
};


/**
 * @param {!Array<!proto.ethermint.evm.v1.State>} value
 * @return {!proto.ethermint.evm.v1.GenesisAccount} returns this
*/
proto.ethermint.evm.v1.GenesisAccount.prototype.setStorageList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 3, value);
};


/**
 * @param {!proto.ethermint.evm.v1.State=} opt_value
 * @param {number=} opt_index
 * @return {!proto.ethermint.evm.v1.State}
 */
proto.ethermint.evm.v1.GenesisAccount.prototype.addStorage = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 3, opt_value, proto.ethermint.evm.v1.State, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.ethermint.evm.v1.GenesisAccount} returns this
 */
proto.ethermint.evm.v1.GenesisAccount.prototype.clearStorageList = function() {
  return this.setStorageList([]);
};


