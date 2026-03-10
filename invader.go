components {
  id: "invader"
  component: "/main/invader.script"
}
components {
  id: "explosion"
  component: "/main/explosion.particlefx"
}
components {
  id: "invader retry"
  component: "/main/invader retry.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"alien\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/graphics.atlas\"\n"
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
  "group: \"default\"\n"
  "mask: \"default\"\n"
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
  "  data: 22.100353\n"
  "  data: 21.613468\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
