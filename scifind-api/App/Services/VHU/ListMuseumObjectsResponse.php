<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: vhu/v1/services.proto

namespace App\Services\VHU;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>vhu.v1.ListMuseumObjectsResponse</code>
 */
class ListMuseumObjectsResponse extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>repeated .vhu.v1.MuseumObject objects = 1 [json_name = "objects"];</code>
     */
    private $objects;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type array<\App\Services\VHU\MuseumObject>|\Google\Protobuf\Internal\RepeatedField $objects
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Vhu\V1\Services::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>repeated .vhu.v1.MuseumObject objects = 1 [json_name = "objects"];</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getObjects()
    {
        return $this->objects;
    }

    /**
     * Generated from protobuf field <code>repeated .vhu.v1.MuseumObject objects = 1 [json_name = "objects"];</code>
     * @param array<\App\Services\VHU\MuseumObject>|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setObjects($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \App\Services\VHU\MuseumObject::class);
        $this->objects = $arr;

        return $this;
    }

}

