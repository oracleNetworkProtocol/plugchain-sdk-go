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

goog.provide('proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement');
goog.provide('proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement.SecurityRequirementValue');

goog.require('jspb.BinaryReader');
goog.require('jspb.BinaryWriter');
goog.require('jspb.Map');
goog.require('jspb.Message');

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
proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement.displayName = 'proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement';
}
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
proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement.SecurityRequirementValue = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement.SecurityRequirementValue.repeatedFields_, null);
};
goog.inherits(proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement.SecurityRequirementValue, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement.SecurityRequirementValue.displayName = 'proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement.SecurityRequirementValue';
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
proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement.prototype.toObject = function(opt_includeInstance) {
  return proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement.toObject = function(includeInstance, msg) {
  var f, obj = {
    securityRequirementMap: (f = msg.getSecurityRequirementMap()) ? f.toObject(includeInstance, proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement.SecurityRequirementValue.toObject) : []
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
 * @return {!proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement}
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement;
  return proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement}
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = msg.getSecurityRequirementMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readString, jspb.BinaryReader.prototype.readMessage, proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement.SecurityRequirementValue.deserializeBinaryFromReader, "", new proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement.SecurityRequirementValue());
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
proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getSecurityRequirementMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(1, writer, jspb.BinaryWriter.prototype.writeString, jspb.BinaryWriter.prototype.writeMessage, proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement.SecurityRequirementValue.serializeBinaryToWriter);
  }
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement.SecurityRequirementValue.repeatedFields_ = [1];



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
proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement.SecurityRequirementValue.prototype.toObject = function(opt_includeInstance) {
  return proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement.SecurityRequirementValue.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement.SecurityRequirementValue} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement.SecurityRequirementValue.toObject = function(includeInstance, msg) {
  var f, obj = {
    scopeList: (f = jspb.Message.getRepeatedField(msg, 1)) == null ? undefined : f
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
 * @return {!proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement.SecurityRequirementValue}
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement.SecurityRequirementValue.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement.SecurityRequirementValue;
  return proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement.SecurityRequirementValue.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement.SecurityRequirementValue} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement.SecurityRequirementValue}
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement.SecurityRequirementValue.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.addScope(value);
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
proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement.SecurityRequirementValue.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement.SecurityRequirementValue.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement.SecurityRequirementValue} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement.SecurityRequirementValue.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getScopeList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      1,
      f
    );
  }
};


/**
 * repeated string scope = 1;
 * @return {!Array<string>}
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement.SecurityRequirementValue.prototype.getScopeList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 1));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement.SecurityRequirementValue} returns this
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement.SecurityRequirementValue.prototype.setScopeList = function(value) {
  return jspb.Message.setField(this, 1, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement.SecurityRequirementValue} returns this
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement.SecurityRequirementValue.prototype.addScope = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 1, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement.SecurityRequirementValue} returns this
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement.SecurityRequirementValue.prototype.clearScopeList = function() {
  return this.setScopeList([]);
};


/**
 * map<string, SecurityRequirementValue> security_requirement = 1;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<string,!proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement.SecurityRequirementValue>}
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement.prototype.getSecurityRequirementMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<string,!proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement.SecurityRequirementValue>} */ (
      jspb.Message.getMapField(this, 1, opt_noLazyCreate,
      proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement.SecurityRequirementValue));
};


/**
 * Clears values from the map. The map will be non-null.
 * @return {!proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement} returns this
 */
proto.grpc.gateway.protoc_gen_openapiv2.options.SecurityRequirement.prototype.clearSecurityRequirementMap = function() {
  this.getSecurityRequirementMap().clear();
  return this;};


