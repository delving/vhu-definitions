<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: vhu/v1/definitions.proto

namespace App\Services\VHU;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>vhu.v1.MuseumObject</code>
 */
class MuseumObject extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>.vhu.v1.Header header = 1 [json_name = "header"];</code>
     */
    protected $header = null;
    /**
     * Generated from protobuf field <code>.vhu.v1.Metadata metadata = 2 [json_name = "metadata"];</code>
     */
    protected $metadata = null;
    /**
     * skos:relatedMatch via thesaurus link
     *
     * Generated from protobuf field <code>repeated .vhu.v1.ThesaurusConcept concepts = 3 [json_name = "concepts"];</code>
     */
    private $concepts;
    /**
     * Generated from protobuf field <code>repeated .vhu.v1.MuseumObject related_objects = 4 [json_name = "relatedObjects"];</code>
     */
    private $related_objects;
    /**
     * Generated from protobuf field <code>.vhu.v1.ObjLinkInfo link_info = 5 [json_name = "linkInfo"];</code>
     */
    protected $link_info = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \App\Services\VHU\Header $header
     *     @type \App\Services\VHU\Metadata $metadata
     *     @type array<\App\Services\VHU\ThesaurusConcept>|\Google\Protobuf\Internal\RepeatedField $concepts
     *           skos:relatedMatch via thesaurus link
     *     @type array<\App\Services\VHU\MuseumObject>|\Google\Protobuf\Internal\RepeatedField $related_objects
     *     @type \App\Services\VHU\ObjLinkInfo $link_info
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
     * Generated from protobuf field <code>.vhu.v1.Metadata metadata = 2 [json_name = "metadata"];</code>
     * @return \App\Services\VHU\Metadata|null
     */
    public function getMetadata()
    {
        return $this->metadata;
    }

    public function hasMetadata()
    {
        return isset($this->metadata);
    }

    public function clearMetadata()
    {
        unset($this->metadata);
    }

    /**
     * Generated from protobuf field <code>.vhu.v1.Metadata metadata = 2 [json_name = "metadata"];</code>
     * @param \App\Services\VHU\Metadata $var
     * @return $this
     */
    public function setMetadata($var)
    {
        GPBUtil::checkMessage($var, \App\Services\VHU\Metadata::class);
        $this->metadata = $var;

        return $this;
    }

    /**
     * skos:relatedMatch via thesaurus link
     *
     * Generated from protobuf field <code>repeated .vhu.v1.ThesaurusConcept concepts = 3 [json_name = "concepts"];</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getConcepts()
    {
        return $this->concepts;
    }

    /**
     * skos:relatedMatch via thesaurus link
     *
     * Generated from protobuf field <code>repeated .vhu.v1.ThesaurusConcept concepts = 3 [json_name = "concepts"];</code>
     * @param array<\App\Services\VHU\ThesaurusConcept>|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setConcepts($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \App\Services\VHU\ThesaurusConcept::class);
        $this->concepts = $arr;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated .vhu.v1.MuseumObject related_objects = 4 [json_name = "relatedObjects"];</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getRelatedObjects()
    {
        return $this->related_objects;
    }

    /**
     * Generated from protobuf field <code>repeated .vhu.v1.MuseumObject related_objects = 4 [json_name = "relatedObjects"];</code>
     * @param array<\App\Services\VHU\MuseumObject>|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setRelatedObjects($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \App\Services\VHU\MuseumObject::class);
        $this->related_objects = $arr;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.vhu.v1.ObjLinkInfo link_info = 5 [json_name = "linkInfo"];</code>
     * @return \App\Services\VHU\ObjLinkInfo|null
     */
    public function getLinkInfo()
    {
        return $this->link_info;
    }

    public function hasLinkInfo()
    {
        return isset($this->link_info);
    }

    public function clearLinkInfo()
    {
        unset($this->link_info);
    }

    /**
     * Generated from protobuf field <code>.vhu.v1.ObjLinkInfo link_info = 5 [json_name = "linkInfo"];</code>
     * @param \App\Services\VHU\ObjLinkInfo $var
     * @return $this
     */
    public function setLinkInfo($var)
    {
        GPBUtil::checkMessage($var, \App\Services\VHU\ObjLinkInfo::class);
        $this->link_info = $var;

        return $this;
    }

}

