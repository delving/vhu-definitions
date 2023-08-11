<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: vhu/v1/definitions.proto

namespace App\Services\VHU;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>vhu.v1.Event</code>
 */
class Event extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>.vhu.v1.Header header = 1 [json_name = "header"];</code>
     */
    protected $header = null;
    /**
     * Generated from protobuf field <code>string start_date = 2 [json_name = "startDate"];</code>
     */
    protected $start_date = '';
    /**
     * Generated from protobuf field <code>string end_date = 3 [json_name = "endDate"];</code>
     */
    protected $end_date = '';
    /**
     * Generated from protobuf field <code>string period = 4 [json_name = "period"];</code>
     */
    protected $period = '';
    /**
     * Generated from protobuf field <code>repeated .vhu.v1.MuseumObjectInfo objects = 5 [json_name = "objects"];</code>
     */
    private $objects;
    /**
     * Generated from protobuf field <code>bool is_group = 6 [json_name = "isGroup"];</code>
     */
    protected $is_group = false;
    /**
     * Generated from protobuf field <code>string year = 7 [json_name = "year"];</code>
     */
    protected $year = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \App\Services\VHU\Header $header
     *     @type string $start_date
     *     @type string $end_date
     *     @type string $period
     *     @type array<\App\Services\VHU\MuseumObjectInfo>|\Google\Protobuf\Internal\RepeatedField $objects
     *     @type bool $is_group
     *     @type string $year
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Vhu\V1\Definitions::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>.vhu.v1.Header header = 1 [json_name = "header"];</code>
     * @return \App\Services\VHU\Header|null
     */
    public function getHeader()
    {
        return $this->header;
    }

    public function hasHeader()
    {
        return isset($this->header);
    }

    public function clearHeader()
    {
        unset($this->header);
    }

    /**
     * Generated from protobuf field <code>.vhu.v1.Header header = 1 [json_name = "header"];</code>
     * @param \App\Services\VHU\Header $var
     * @return $this
     */
    public function setHeader($var)
    {
        GPBUtil::checkMessage($var, \App\Services\VHU\Header::class);
        $this->header = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string start_date = 2 [json_name = "startDate"];</code>
     * @return string
     */
    public function getStartDate()
    {
        return $this->start_date;
    }

    /**
     * Generated from protobuf field <code>string start_date = 2 [json_name = "startDate"];</code>
     * @param string $var
     * @return $this
     */
    public function setStartDate($var)
    {
        GPBUtil::checkString($var, True);
        $this->start_date = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string end_date = 3 [json_name = "endDate"];</code>
     * @return string
     */
    public function getEndDate()
    {
        return $this->end_date;
    }

    /**
     * Generated from protobuf field <code>string end_date = 3 [json_name = "endDate"];</code>
     * @param string $var
     * @return $this
     */
    public function setEndDate($var)
    {
        GPBUtil::checkString($var, True);
        $this->end_date = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string period = 4 [json_name = "period"];</code>
     * @return string
     */
    public function getPeriod()
    {
        return $this->period;
    }

    /**
     * Generated from protobuf field <code>string period = 4 [json_name = "period"];</code>
     * @param string $var
     * @return $this
     */
    public function setPeriod($var)
    {
        GPBUtil::checkString($var, True);
        $this->period = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated .vhu.v1.MuseumObjectInfo objects = 5 [json_name = "objects"];</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getObjects()
    {
        return $this->objects;
    }

    /**
     * Generated from protobuf field <code>repeated .vhu.v1.MuseumObjectInfo objects = 5 [json_name = "objects"];</code>
     * @param array<\App\Services\VHU\MuseumObjectInfo>|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setObjects($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \App\Services\VHU\MuseumObjectInfo::class);
        $this->objects = $arr;

        return $this;
    }

    /**
     * Generated from protobuf field <code>bool is_group = 6 [json_name = "isGroup"];</code>
     * @return bool
     */
    public function getIsGroup()
    {
        return $this->is_group;
    }

    /**
     * Generated from protobuf field <code>bool is_group = 6 [json_name = "isGroup"];</code>
     * @param bool $var
     * @return $this
     */
    public function setIsGroup($var)
    {
        GPBUtil::checkBool($var);
        $this->is_group = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string year = 7 [json_name = "year"];</code>
     * @return string
     */
    public function getYear()
    {
        return $this->year;
    }

    /**
     * Generated from protobuf field <code>string year = 7 [json_name = "year"];</code>
     * @param string $var
     * @return $this
     */
    public function setYear($var)
    {
        GPBUtil::checkString($var, True);
        $this->year = $var;

        return $this;
    }

}

