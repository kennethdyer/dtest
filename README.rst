#####
DTest
#####

DTest is a tool to test multi-platform deployment instructions
in Podman before the content ships to public documentation
sites.  Instructions are stored in YAML and sync to
reStructuredText files.

DTest is intended for writers who *detest* when installation
pages get fouled by untested instructions. 

Installation
************

To install DTest, use the Go command:

.. code-block:: console

   go install github.com/kennethdyer/dtest

Usage
*****

TODO

History
*******

v0.1.x
======

Status: *In-progress*

The target for this release series is a proof of concept for
deployment testing.  We aim to implement the following features:

* [TODO]: Implement `Cobra <https://github.com/spf13/cobra>`_ and 
  `Viper <https://github.com/spf13/viper>`_ interfaces for CLI.
* [TODO]: Implement a logging module, with pass through to `Log
  <https://github.com/charmbracelet/log>`_ and level
  configuration.
* [TODO]: Implement functions and commands to configure, status start,
  and stop the Podman machine.  Add a conditional call to the stop
  function to the shutdown process in the logging module.
* [TODO]: Implement image registry and configuration.  Add a list
  command to show what images are available to the tool.

_Note_, there is no operation to provision containers at this
stage. Deployments will be tested against manually built
containers.
