/** Samples for Controllers List. */
public final class Main {
    /*
     * x-ms-original-file: specification/devspaces/resource-manager/Microsoft.DevSpaces/stable/2019-04-01/examples/ControllersList_example.json
     */
    /**
     * Sample code: ControllersList.
     *
     * @param manager Entry point to DevSpacesManager.
     */
    public static void controllersList(com.azure.resourcemanager.devspaces.DevSpacesManager manager) {
        manager.controllers().list(com.azure.core.util.Context.NONE);
    }
}
