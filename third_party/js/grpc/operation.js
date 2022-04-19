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

goog.provide('proto.grpc.gateway.protoc_gen_openapiv2.options.Operation');

goog.require('jspb.BinaryReader');
goog.require('jspb.BinaryWriter');
goog.require('jspb.Map');
goog.require('jspb.Message');
goog.require('proto.google.protobuf.Value');
goog.require('proto.grpc.gateway.protoc_gen_openapiv2.options.ExternalDocumentation');
goog.require('proto.grpc.gateway.protoc_gen_openapiv2.options.Response');
goog.require('proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement');

goog.forwardDeclare('proto.grpc.gateway.protoc_gen_openapiv2.options.Scheme');
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
proto.grpc.gateway.protoc_gen_openapiv2.options.Operation = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.grpc.gateway.protoc_gen_openapiv2.options.Operation.repeatedFields_, null);
};
goog.inherits(proto.grpc.gateway.protoc_gen_openapiv2.options.Operation, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.grpc.gateway.protoc_gen_openapiv2.options.Operation.displayName = 'proto.grpc.gateway.protoc_gen_openapiv2.options.Operation';
}

/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.Operation.repeatedFields_ = [1,6,7,10,12];



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
proto.grpc.gateway.protoc_gen_openapiv2.options.Operation.prototype.toObject = function(opt_includeInstance) {
  return proto.grpc.gateway.protoc_gen_openapiv2.options.Operation.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.grpc.gateway.protoc_gen_openapiv2.options.Operation} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.Operation.toObject = function(includeInstance, msg) {
  var f, obj = {
    tagsList: (f = jspb.Message.getRepeatedField(msg, 1)) == null ? undefined : f,
    summary: jspb.Message.getFieldWithDefault(msg, 2, ""),
    description: jspb.Message.getFieldWithDefault(msg, 3, ""),
    externalDocs: (f = msg.getExternalDocs()) && proto.grpc.gateway.protoc_gen_openapiv2.options.ExternalDocumentation.toObject(includeInstance, f),
    operationId: jspb.Message.getFieldWithDefault(msg, 5, ""),
    consumesList: (f = jspb.Message.getRepeatedField(msg, 6)) == null ? undefined : f,
    producesList: (f = jspb.Message.getRepeatedField(msg, 7)) == null ? undefined : f,
    responsesMap: (f = msg.getResponsesMap()) ? f.toObject(includeInstance, proto.grpc.gateway.protoc_gen_openapiv2.options.Response.toObject) : [],
    schemesList: (f = jspb.Message.getRepeatedField(msg, 10)) == null ? undefined : f,
    deprecated: jspb.Message.getBooleanFieldWithDefault(msg, 11, false),
    securityList: jspb.Message.toObjectList(msg.getSecurityList(),
    proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement.toObject, includeInstance),
    extensionsMap: (f = msg.getExtensionsMap()) ? f.toObject(includeInstance, proto.google.protobuf.Value.toObject) : []
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
 * @return {!proto.grpc.gateway.protoc_gen_openapiv2.options.Operation}
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.Operation.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.grpc.gateway.protoc_gen_openapiv2.options.Operation;
  return proto.grpc.gateway.protoc_gen_openapiv2.options.Operation.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.grpc.gateway.protoc_gen_openapiv2.options.Operation} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.grpc.gateway.protoc_gen_openapiv2.options.Operation}
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.Operation.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.addTags(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setSummary(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setDescription(value);
      break;
    case 4:
      var value = new proto.grpc.gateway.protoc_gen_openapiv2.options.ExternalDocumentation;
      reader.readMessage(value,proto.grpc.gateway.protoc_gen_openapiv2.options.ExternalDocumentation.deserializeBinaryFromReader);
      msg.setExternalDocs(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setOperationId(value);
      break;
    case 6:
      var value = /** @type {string} */ (reader.readString());
      msg.addConsumes(value);
      break;
    case 7:
      var value = /** @type {string} */ (reader.readString());
      msg.addProduces(value);
      break;
    case 9:
      var value = msg.getResponsesMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readString, jspb.BinaryReader.prototype.readMessage, proto.grpc.gateway.protoc_gen_openapiv2.options.Response.deserializeBinaryFromReader, "", new proto.grpc.gateway.protoc_gen_openapiv2.options.Response());
         });
      break;
    case 10:
      var values = /** @type {!Array<!proto.grpc.gateway.protoc_gen_openapiv2.options.Scheme>} */ (reader.isDelimited() ? reader.readPackedEnum() : [reader.readEnum()]);
      for (var i = 0; i < values.length; i++) {
        msg.addSchemes(values[i]);
      }
      break;
    case 11:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setDeprecated(value);
      break;
    case 12:
      var value = new proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement;
      reader.readMessage(value,proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement.deserializeBinaryFromReader);
      msg.addSecurity(value);
      break;
    case 13:
      var value = msg.getExtensionsMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readString, jspb.BinaryReader.prototype.readMessage, proto.google.protobuf.Value.deserializeBinaryFromReader, "", new proto.google.protobuf.Value());
         });
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
proto.grpc.gateway.protoc_gen_openapiv2.options.Operation.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.grpc.gateway.protoc_gen_openapiv2.options.Operation.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.grpc.gateway.protoc_gen_openapiv2.options.Operation} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.Operation.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getTagsList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      1,
      f
    );
  }
  f = message.getSummary();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getDescription();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getExternalDocs();
  if (f != null) {
    writer.writeMessage(
      4,
      f,
      proto.grpc.gateway.protoc_gen_openapiv2.options.ExternalDocumentation.serializeBinaryToWriter
    );
  }
  f = message.getOperationId();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
  f = message.getConsumesList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      6,
      f
    );
  }
  f = message.getProducesList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      7,
      f
    );
  }
  f = message.getResponsesMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(9, writer, jspb.BinaryWriter.prototype.writeString, jspb.BinaryWriter.prototype.writeMessage, proto.grpc.gateway.protoc_gen_openapiv2.options.Response.serializeBinaryToWriter);
  }
  f = message.getSchemesList();
  if (f.length > 0) {
    writer.writePackedEnum(
      10,
      f
    );
  }
  f = message.getDeprecated();
  if (f) {
    writer.writeBool(
      11,
      f
    );
  }
  f = message.getSecurityList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      12,
      f,
      proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement.serializeBinaryToWriter
    );
  }
  f = message.getExtensionsMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(13, writer, jspb.BinaryWriter.prototype.writeString, jspb.BinaryWriter.prototype.writeMessage, proto.google.protobuf.Value.serializeBinaryToWriter);
  }
};


/**
 * repeated string tags = 1;
 * @return {!Array<string>}
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.Operation.prototype.getTagsList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 1));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.grpc.gateway.protoc_gen_openapiv2.options.Operation} returns this
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.Operation.prototype.setTagsList = function(value) {
  return jspb.Message.setField(this, 1, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.grpc.gateway.protoc_gen_openapiv2.options.Operation} returns this
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.Operation.prototype.addTags = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 1, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.grpc.gateway.protoc_gen_openapiv2.options.Operation} returns this
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.Operation.prototype.clearTagsList = function() {
  return this.setTagsList([]);
};


/**
 * optional string summary = 2;
 * @return {string}
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.Operation.prototype.getSummary = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.grpc.gateway.protoc_gen_openapiv2.options.Operation} returns this
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.Operation.prototype.setSummary = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string description = 3;
 * @return {string}
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.Operation.prototype.getDescription = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.grpc.gateway.protoc_gen_openapiv2.options.Operation} returns this
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.Operation.prototype.setDescription = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional ExternalDocumentation external_docs = 4;
 * @return {?proto.grpc.gateway.protoc_gen_openapiv2.options.ExternalDocumentation}
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.Operation.prototype.getExternalDocs = function() {
  return /** @type{?proto.grpc.gateway.protoc_gen_openapiv2.options.ExternalDocumentation} */ (
    jspb.Message.getWrapperField(this, proto.grpc.gateway.protoc_gen_openapiv2.options.ExternalDocumentation, 4));
};


/**
 * @param {?proto.grpc.gateway.protoc_gen_openapiv2.options.ExternalDocumentation|undefined} value
 * @return {!proto.grpc.gateway.protoc_gen_openapiv2.options.Operation} returns this
*/
proto.grpc.gateway.protoc_gen_openapiv2.options.Operation.prototype.setExternalDocs = function(value) {
  return jspb.Message.setWrapperField(this, 4, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.grpc.gateway.protoc_gen_openapiv2.options.Operation} returns this
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.Operation.prototype.clearExternalDocs = function() {
  return this.setExternalDocs(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.Operation.prototype.hasExternalDocs = function() {
  return jspb.Message.getField(this, 4) != null;
};


/**
 * optional string operation_id = 5;
 * @return {string}
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.Operation.prototype.getOperationId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/**
 * @param {string} value
 * @return {!proto.grpc.gateway.protoc_gen_openapiv2.options.Operation} returns this
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.Operation.prototype.setOperationId = function(value) {
  return jspb.Message.setProto3StringField(this, 5, value);
};


/**
 * repeated string consumes = 6;
 * @return {!Array<string>}
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.Operation.prototype.getConsumesList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 6));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.grpc.gateway.protoc_gen_openapiv2.options.Operation} returns this
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.Operation.prototype.setConsumesList = function(value) {
  return jspb.Message.setField(this, 6, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.grpc.gateway.protoc_gen_openapiv2.options.Operation} returns this
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.Operation.prototype.addConsumes = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 6, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.grpc.gateway.protoc_gen_openapiv2.options.Operation} returns this
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.Operation.prototype.clearConsumesList = function() {
  return this.setConsumesList([]);
};


/**
 * repeated string produces = 7;
 * @return {!Array<string>}
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.Operation.prototype.getProducesList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 7));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.grpc.gateway.protoc_gen_openapiv2.options.Operation} returns this
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.Operation.prototype.setProducesList = function(value) {
  return jspb.Message.setField(this, 7, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.grpc.gateway.protoc_gen_openapiv2.options.Operation} returns this
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.Operation.prototype.addProduces = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 7, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.grpc.gateway.protoc_gen_openapiv2.options.Operation} returns this
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.Operation.prototype.clearProducesList = function() {
  return this.setProducesList([]);
};


/**
 * map<string, Response> responses = 9;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<string,!proto.grpc.gateway.protoc_gen_openapiv2.options.Response>}
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.Operation.prototype.getResponsesMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<string,!proto.grpc.gateway.protoc_gen_openapiv2.options.Response>} */ (
      jspb.Message.getMapField(this, 9, opt_noLazyCreate,
      proto.grpc.gateway.protoc_gen_openapiv2.options.Response));
};


/**
 * Clears values from the map. The map will be non-null.
 * @return {!proto.grpc.gateway.protoc_gen_openapiv2.options.Operation} returns this
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.Operation.prototype.clearResponsesMap = function() {
  this.getResponsesMap().clear();
  return this;};


/**
 * repeated Scheme schemes = 10;
 * @return {!Array<!proto.grpc.gateway.protoc_gen_openapiv2.options.Scheme>}
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.Operation.prototype.getSchemesList = function() {
  return /** @type {!Array<!proto.grpc.gateway.protoc_gen_openapiv2.options.Scheme>} */ (jspb.Message.getRepeatedField(this, 10));
};


/**
 * @param {!Array<!proto.grpc.gateway.protoc_gen_openapiv2.options.Scheme>} value
 * @return {!proto.grpc.gateway.protoc_gen_openapiv2.options.Operation} returns this
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.Operation.prototype.setSchemesList = function(value) {
  return jspb.Message.setField(this, 10, value || []);
};


/**
 * @param {!proto.grpc.gateway.protoc_gen_openapiv2.options.Scheme} value
 * @param {number=} opt_index
 * @return {!proto.grpc.gateway.protoc_gen_openapiv2.options.Operation} returns this
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.Operation.prototype.addSchemes = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 10, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.grpc.gateway.protoc_gen_openapiv2.options.Operation} returns this
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.Operation.prototype.clearSchemesList = function() {
  return this.setSchemesList([]);
};


/**
 * optional bool deprecated = 11;
 * @return {boolean}
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.Operation.prototype.getDeprecated = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 11, false));
};


/**
 * @param {boolean} value
 * @return {!proto.grpc.gateway.protoc_gen_openapiv2.options.Operation} returns this
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.Operation.prototype.setDeprecated = function(value) {
  return jspb.Message.setProto3BooleanField(this, 11, value);
};


/**
 * repeated SecurityRequirement security = 12;
 * @return {!Array<!proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement>}
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.Operation.prototype.getSecurityList = function() {
  return /** @type{!Array<!proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement, 12));
};


/**
 * @param {!Array<!proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement>} value
 * @return {!proto.grpc.gateway.protoc_gen_openapiv2.options.Operation} returns this
*/
proto.grpc.gateway.protoc_gen_openapiv2.options.Operation.prototype.setSecurityList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 12, value);
};


/**
 * @param {!proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement=} opt_value
 * @param {number=} opt_index
 * @return {!proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement}
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.Operation.prototype.addSecurity = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 12, opt_value, proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.grpc.gateway.protoc_gen_openapiv2.options.Operation} returns this
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.Operation.prototype.clearSecurityList = function() {
  return this.setSecurityList([]);
};


/**
 * map<string, google.protobuf.Value> extensions = 13;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<string,!proto.google.protobuf.Value>}
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.Operation.prototype.getExtensionsMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<string,!proto.google.protobuf.Value>} */ (
      jspb.Message.getMapField(this, 13, opt_noLazyCreate,
      proto.google.protobuf.Value));
};


/**
 * Clears values from the map. The map will be non-null.
 * @return {!proto.grpc.gateway.protoc_gen_openapiv2.options.Operation} returns this
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.Operation.prototype.clearExtensionsMap = function() {
  this.getExtensionsMap().clear();
  return this;};


