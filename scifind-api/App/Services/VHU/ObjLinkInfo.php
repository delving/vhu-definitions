<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: vhu/v1/definitions.proto

namespace App\Services\VHU;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>vhu.v1.ObjLinkInfo</code>
 */
class ObjLinkInfo extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string museum_name = 1 [json_name = "museumName"];</code>
     */
    protected $museum_name = '';
    /**
     * Generated from protobuf field <code>string object_number = 2 [json_name = "objectNumber"];</code>
     */
    protected $object_number = '';
    /**
     * Generated from protobuf field <code>string object_name = 3 [json_name = "objectName"];</code>
     */
    protected $object_name = '';
    /**
     * Generated from protobuf field <code>string dcn_uri = 4 [json_name = "dcnUri"];</code>
     */
    protected $dcn_uri = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $museum_name
     *     @type string $object_number
     *     @type string $object_name
     *     @type string $dcn_uri
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Vhu\V1\Definitions::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string museum_name = 1 [json_name = "museumName"];</code>
     * @return string
     */
    public function getMuseumName()
    {
        return $this->museum_name;
    }

    /**
     * Generated from protobuf field <code>string museum_name = 1 [json_name = "museumName"];</code>
     * @param string $var
     * @return $this
     */
    public function setMuseumName($var)
    {
        GPBUtil::checkString($var, True);
        $this->museum_name = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string object_number = 2 [json_name = "objectNumber"];</code>
     * @return string
     */
    public function getObjectNumber()
    {
        return $this->object_number;
    }

    /**
     * Generated from protobuf field <code>string object_number = 2 [json_name = "objectNumber"];</code>
     * @param string $var
     * @return $this
     */
    public function setObjectNumber($var)
    {
        GPBUtil::checkString($var, True);
        $this->object_number = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string object_name = 3 [json_name = "objectName"];</code>
     * @return string
     */
    public function getObjectName()
    {
        return $this->object_name;
    }

    /**
     * Generated from protobuf field <code>string object_name = 3 [json_name = "objectName"];</code>
     * @param string $var
     * @return $this
     */
    public function setObjectName($var)
    {
        GPBUtil::checkString($var, True);
        $this->object_name = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string dcn_uri = 4 [json_name = "dcnUri"];</code>
     * @return string
     */
    public function getDcnUri()
    {
        return $this->dcn_uri;
    }

    /**
     * Generated from protobuf field <code>string dcn_uri = 4 [json_name = "dcnUri"];</code>
     * @param string $var
     * @return $this
     */
    public function setDcnUri($var)
    {
        GPBUtil::checkString($var, True);
        $this->dcn_uri = $var;

        return $this;
    }

}

