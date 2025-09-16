components {
  id: "screen_proxy"
  component: "/monarch/screen_proxy.script"
  properties {
    id: "screen_id"
    value: "mainmenu"
    type: PROPERTY_TYPE_HASH
  }
}
embedded_components {
  id: "collectionproxy"
  type: "collectionproxy"
  data: "collection: \"/game/mainmenu.collection\"\n"
}
