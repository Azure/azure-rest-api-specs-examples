/** Samples for Controllers ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/devspaces/resource-manager/Microsoft.DevSpaces/stable/2019-04-01/examples/ControllersListByResourceGroup_example.json
     */
    /**
     * Sample code: ControllersListByResourceGroup.
     *
     * @param manager Entry point to DevSpacesManager.
     */
    public static void controllersListByResourceGroup(com.azure.resourcemanager.devspaces.DevSpacesManager manager) {
        manager.controllers().listByResourceGroup("myResourceGroup", com.azure.core.util.Context.NONE);
    }
}
