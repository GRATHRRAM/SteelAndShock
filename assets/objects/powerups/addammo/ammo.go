components {
  id: "script"
  component: "/assets/objects/powerups/PowerUpAnim.script"
  properties {
    id: "powerup"
    value: "ammo"
    type: PROPERTY_TYPE_HASH
  }
  properties {
    id: "val"
    value: "5.0"
    type: PROPERTY_TYPE_NUMBER
  }
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"ammo\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/objects/powerups/PowerUpsTxt.atlas\"\n"
  "}\n"
  ""
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"powerup\"\n"
  "mask: \"player\"\n"
  "mask: \"player2\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 100.327866\n"
  "  data: 49.696968\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
