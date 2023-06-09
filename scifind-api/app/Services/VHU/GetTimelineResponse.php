<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: vhu/v1/services.proto

namespace App\Services\VHU;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>vhu.v1.GetTimelineResponse</code>
 */
class GetTimelineResponse extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>.vhu.v1.Timeline object = 1 [json_name = "object"];</code>
     */
    protected $object = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \App\Services\VHU\Timeline $object
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Vhu\V1\Services::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>.vhu.v1.Timeline object = 1 [json_name = "object"];</code>
     * @return \App\Services\VHU\Timeline|null
     */
    public function getObject()
    {
        return $this->object;
    }

    public function hasObject()
    {
        return isset($this->object);
    }

    public function clearObject()
    {
        unset($this->object);
    }

    /**
     * Generated from protobuf field <code>.vhu.v1.Timeline object = 1 [json_name = "object"];</code>
     * @param \App\Services\VHU\Timeline $var
     * @return $this
     */
    public function setObject($var)
    {
        GPBUtil::checkMessage($var, \App\Services\VHU\Timeline::class);
        $this->object = $var;

        return $this;
    }

}

