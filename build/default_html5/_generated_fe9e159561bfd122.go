components {
  id: "initloadscreen"
  component: "/game/loadscreen/initloadscreen.script"
}
components {
  id: "testscriptforrive"
  component: "/game/loadscreen/testscriptforrive.script"
}
embedded_components {
  id: "rivemodel"
  type: "rivemodel"
  data: "scene: \"/game/loadscreen/testscene.rivescene\"\ndefault_animation: \"Timeline 1\"\nmaterial: \"/defold-rive/assets/rivemodel.material\"\ndefault_state_machine: \"State Machine 1\"\nartboard: \"Artboard\"\ncoordinate_system: COORDINATE_SYSTEM_RIVE\n"
}
