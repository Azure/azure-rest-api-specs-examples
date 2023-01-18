/** Samples for Controllers Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/devspaces/resource-manager/Microsoft.DevSpaces/stable/2019-04-01/examples/ControllersDelete_example.json
     */
    /**
     * Sample code: ControllersDelete.
     *
     * @param manager Entry point to DevSpacesManager.
     */
    public static void controllersDelete(com.azure.resourcemanager.devspaces.DevSpacesManager manager) {
        manager.controllers().delete("myResourceGroup", "myControllerResource", com.azure.core.util.Context.NONE);
    }
}
