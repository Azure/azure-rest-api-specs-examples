/** Samples for Controllers GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/devspaces/resource-manager/Microsoft.DevSpaces/stable/2019-04-01/examples/ControllersGet_example.json
     */
    /**
     * Sample code: ControllersGet.
     *
     * @param manager Entry point to DevSpacesManager.
     */
    public static void controllersGet(com.azure.resourcemanager.devspaces.DevSpacesManager manager) {
        manager
            .controllers()
            .getByResourceGroupWithResponse(
                "myResourceGroup", "myControllerResource", com.azure.core.util.Context.NONE);
    }
}
