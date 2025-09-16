components {
  id: "screen_proxy"
  component: "/monarch/screen_proxy.script"
  properties {
    id: "screen_id"
    value: "loadscreen"
    type: PROPERTY_TYPE_HASH
  }
}
embedded_components {
  id: "collectionproxy"
  type: "collectionproxy"
  data: "collection: \"/game/loadscreen/loadscreen.collection\"\n"
}
