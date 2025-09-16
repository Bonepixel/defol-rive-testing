components {
  id: "player"
  component: "/game/player.script"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
  property_decls {
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "collision_shape: \"\"\ntype: COLLISION_OBJECT_TYPE_KINEMATIC\nmass: 0.0\nfriction: 0.1\nrestitution: 0.5\ngroup: \"player\"\nmask: \"ground\"\nmask: \"danger\"\nembedded_collision_shape {\n  shapes {\n    shape_type: TYPE_SPHERE\n    position {\n      x: 0.0\n      y: 40.0\n      z: 0.0\n    }\n    rotation {\n      x: 0.0\n      y: 0.0\n      z: 0.0\n      w: 1.0\n    }\n    index: 0\n    count: 1\n  }\n  shapes {\n    shape_type: TYPE_SPHERE\n    position {\n      x: 0.0\n      y: 100.0\n      z: 0.0\n    }\n    rotation {\n      x: 0.0\n      y: 0.0\n      z: 0.0\n      w: 1.0\n    }\n    index: 1\n    count: 1\n  }\n  data: 40.0\n  data: 45.0\n}\nlinear_damping: 0.0\nangular_damping: 0.0\nlocked_rotation: false\nbullet: false\n"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "tile_set: \"/assets/player.atlas\"\ndefault_animation: \"walk\"\nmaterial: \"/builtins/materials/sprite.material\"\nblend_mode: BLEND_MODE_ALPHA\n"
  position {
    x: 0.0
    y: 128.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "camera"
  type: "camera"
  data: "aspect_ratio: 1.0\nfov: 0.7854\nnear_z: 1.0\nfar_z: -1.0\nauto_aspect_ratio: 0\northographic_projection: 1\northographic_zoom: 1.0\n"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
