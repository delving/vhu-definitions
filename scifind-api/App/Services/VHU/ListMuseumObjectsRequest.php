<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: vhu/v1/services.proto

namespace App\Services\VHU;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>vhu.v1.ListMuseumObjectsRequest</code>
 */
class ListMuseumObjectsRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string query = 1 [json_name = "query"];</code>
     */
    protected $query = '';
    /**
     * Generated from protobuf field <code>int32 page = 2 [json_name = "page"];</code>
     */
    protected $page = 0;
    /**
     * Generated from protobuf field <code>repeated string query_filters = 3 [json_name = "queryFilters"];</code>
     */
    private $query_filters;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $query
     *     @type int $page
     *     @type array<string>|\Google\Protobuf\Internal\RepeatedField $query_filters
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Vhu\V1\Services::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string query = 1 [json_name = "query"];</code>
     * @return string
     */
    public function getQuery()
    {
        return $this->query;
    }

    /**
     * Generated from protobuf field <code>string query = 1 [json_name = "query"];</code>
     * @param string $var
     * @return $this
     */
    public function setQuery($var)
    {
        GPBUtil::checkString($var, True);
        $this->query = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 page = 2 [json_name = "page"];</code>
     * @return int
     */
    public function getPage()
    {
        return $this->page;
    }

    /**
     * Generated from protobuf field <code>int32 page = 2 [json_name = "page"];</code>
     * @param int $var
     * @return $this
     */
    public function setPage($var)
    {
        GPBUtil::checkInt32($var);
        $this->page = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated string query_filters = 3 [json_name = "queryFilters"];</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getQueryFilters()
    {
        return $this->query_filters;
    }

    /**
     * Generated from protobuf field <code>repeated string query_filters = 3 [json_name = "queryFilters"];</code>
     * @param array<string>|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setQueryFilters($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::STRING);
        $this->query_filters = $arr;

        return $this;
    }

}

