<?php

// autoload_static.php @generated by Composer

namespace Composer\Autoload;

class ComposerStaticInit2ebddedc53dcc6e4f232ecda53abc575
{
    public static $fallbackDirsPsr4 = array (
        0 => __DIR__ . '/../..' . '/Pb',
    );

    public static function getInitializer(ClassLoader $loader)
    {
        return \Closure::bind(function () use ($loader) {
            $loader->fallbackDirsPsr4 = ComposerStaticInit2ebddedc53dcc6e4f232ecda53abc575::$fallbackDirsPsr4;

        }, null, ClassLoader::class);
    }
}