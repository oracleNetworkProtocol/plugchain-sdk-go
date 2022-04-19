// source: protoc-gen-openapiv2/options/openapiv2.proto
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

goog.provide('proto.grpc.gateway.protoc_gen_openapiv2.options.Schema');

goog.require('jspb.BinaryReader');
goog.require('jspb.BinaryWriter');
goog.require('jspb.Message');
goog.require('proto.grpc.gateway.protoc_gen_openapiv2.options.ExternalDocumentation');
goog.require('proto.grpc.gateway.protoc_gen_openapiv2.options.JSONSchema');

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
proto.grpc.gateway.protoc_gen_openapiv2.options.Schema = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.grpc.gateway.protoc_gen_openapiv2.options.Schema, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.grpc.gateway.protoc_gen_openapiv2.options.Schema.displayName = 'proto.grpc.gateway.protoc_gen_openapiv2.options.Schema';
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
proto.grpc.gateway.protoc_gen_openapiv2.options.Schema.prototype.toObject = function(opt_includeInstance) {
  return proto.grpc.gateway.protoc_gen_openapiv2.options.Schema.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.grpc.gateway.protoc_gen_openapiv2.options.Schema} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.Schema.toObject = function(includeInstance, msg) {
  var f, obj = {
    jsonSchema: (f = msg.getJsonSchema()) && proto.grpc.gateway.protoc_gen_openapiv2.options.JSONSchema.toObject(includeInstance, f),
    discriminator: jspb.Message.getFieldWithDefault(msg, 2, ""),
    readOnly: jspb.Message.getBooleanFieldWithDefault(msg, 3, false),
    externalDocs: (f = msg.getExternalDocs()) && proto.grpc.gateway.protoc_gen_openapiv2.options.ExternalDocumentation.toObject(includeInstance, f),
    example: jspb.Message.getFieldWithDefault(msg, 6, "")
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
 * @return {!proto.grpc.gateway.protoc_gen_openapiv2.options.Schema}
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.Schema.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.grpc.gateway.protoc_gen_openapiv2.options.Schema;
  return proto.grpc.gateway.protoc_gen_openapiv2.options.Schema.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.grpc.gateway.protoc_gen_openapiv2.options.Schema} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.grpc.gateway.protoc_gen_openapiv2.options.Schema}
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.Schema.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.grpc.gateway.protoc_gen_openapiv2.options.JSONSchema;
      reader.readMessage(value,proto.grpc.gateway.protoc_gen_openapiv2.options.JSONSchema.deserializeBinaryFromReader);
      msg.setJsonSchema(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setDiscriminator(value);
      break;
    case 3:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setReadOnly(value);
      break;
    case 5:
      var value = new proto.grpc.gateway.protoc_gen_openapiv2.options.ExternalDocumentation;
      reader.readMessage(value,proto.grpc.gateway.protoc_gen_openapiv2.options.ExternalDocumentation.deserializeBinaryFromReader);
      msg.setExternalDocs(value);
      break;
    case 6:
      var value = /** @type {string} */ (reader.readString());
      msg.setExample(value);
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
proto.grpc.gateway.protoc_gen_openapiv2.options.Schema.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.grpc.gateway.protoc_gen_openapiv2.options.Schema.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.grpc.gateway.protoc_gen_openapiv2.options.Schema} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.Schema.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getJsonSchema();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.grpc.gateway.protoc_gen_openapiv2.options.JSONSchema.serializeBinaryToWriter
    );
  }
  f = message.getDiscriminator();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getReadOnly();
  if (f) {
    writer.writeBool(
      3,
      f
    );
  }
  f = message.getExternalDocs();
  if (f != null) {
    writer.writeMessage(
      5,
      f,
      proto.grpc.gateway.protoc_gen_openapiv2.options.ExternalDocumentation.serializeBinaryToWriter
    );
  }
  f = message.getExample();
  if (f.length > 0) {
    writer.writeString(
      6,
      f
    );
  }
};


/**
 * optional JSONSchema json_schema = 1;
 * @return {?proto.grpc.gateway.protoc_gen_openapiv2.options.JSONSchema}
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.Schema.prototype.getJsonSchema = function() {
  return /** @type{?proto.grpc.gateway.protoc_gen_openapiv2.options.JSONSchema} */ (
    jspb.Message.getWrapperField(this, proto.grpc.gateway.protoc_gen_openapiv2.options.JSONSchema, 1));
};


/**
 * @param {?proto.grpc.gateway.protoc_gen_openapiv2.options.JSONSchema|undefined} value
 * @return {!proto.grpc.gateway.protoc_gen_openapiv2.options.Schema} returns this
*/
proto.grpc.gateway.protoc_gen_openapiv2.options.Schema.prototype.setJsonSchema = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.grpc.gateway.protoc_gen_openapiv2.options.Schema} returns this
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.Schema.prototype.clearJsonSchema = function() {
  return this.setJsonSchema(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.Schema.prototype.hasJsonSchema = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional string discriminator = 2;
 * @return {string}
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.Schema.prototype.getDiscriminator = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.grpc.gateway.protoc_gen_openapiv2.options.Schema} returns this
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.Schema.prototype.setDiscriminator = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional bool read_only = 3;
 * @return {boolean}
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.Schema.prototype.getReadOnly = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 3, false));
};


/**
 * @param {boolean} value
 * @return {!proto.grpc.gateway.protoc_gen_openapiv2.options.Schema} returns this
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.Schema.prototype.setReadOnly = function(value) {
  return jspb.Message.setProto3BooleanField(this, 3, value);
};


/**
 * optional ExternalDocumentation external_docs = 5;
 * @return {?proto.grpc.gateway.protoc_gen_openapiv2.options.ExternalDocumentation}
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.Schema.prototype.getExternalDocs = function() {
  return /** @type{?proto.grpc.gateway.protoc_gen_openapiv2.options.ExternalDocumentation} */ (
    jspb.Message.getWrapperField(this, proto.grpc.gateway.protoc_gen_openapiv2.options.ExternalDocumentation, 5));
};


/**
 * @param {?proto.grpc.gateway.protoc_gen_openapiv2.options.ExternalDocumentation|undefined} value
 * @return {!proto.grpc.gateway.protoc_gen_openapiv2.options.Schema} returns this
*/
proto.grpc.gateway.protoc_gen_openapiv2.options.Schema.prototype.setExternalDocs = function(value) {
  return jspb.Message.setWrapperField(this, 5, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.grpc.gateway.protoc_gen_openapiv2.options.Schema} returns this
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.Schema.prototype.clearExternalDocs = function() {
  return this.setExternalDocs(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.Schema.prototype.hasExternalDocs = function() {
  return jspb.Message.getField(this, 5) != null;
};


/**
 * optional string example = 6;
 * @return {string}
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.Schema.prototype.getExample = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 6, ""));
};


/**
 * @param {string} value
 * @return {!proto.grpc.gateway.protoc_gen_openapiv2.options.Schema} returns this
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.Schema.prototype.setExample = function(value) {
  return jspb.Message.setProto3StringField(this, 6, value);
};


