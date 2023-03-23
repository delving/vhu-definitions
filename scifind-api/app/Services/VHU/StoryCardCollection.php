<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: vhu/v1/definitions.proto

namespace App\Services\VHU;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>vhu.v1.StoryCardCollection</code>
 */
class StoryCardCollection extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>.vhu.v1.Header header = 1 [json_name = "header"];</code>
     */
    protected $header = null;
    /**
     * Generated from protobuf field <code>repeated .vhu.v1.StoryCard story_card = 2 [json_name = "storyCard"];</code>
     */
    private $story_card;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \App\Services\VHU\Header $header
     *     @type array<\App\Services\VHU\StoryCard>|\Google\Protobuf\Internal\RepeatedField $story_card
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
     * Generated from protobuf field <code>repeated .vhu.v1.StoryCard story_card = 2 [json_name = "storyCard"];</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getStoryCard()
    {
        return $this->story_card;
    }

    /**
     * Generated from protobuf field <code>repeated .vhu.v1.StoryCard story_card = 2 [json_name = "storyCard"];</code>
     * @param array<\App\Services\VHU\StoryCard>|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setStoryCard($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \App\Services\VHU\StoryCard::class);
        $this->story_card = $arr;

        return $this;
    }

}
