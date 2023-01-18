import com.azure.resourcemanager.devspaces.models.ListConnectionDetailsParameters;

/** Samples for Controllers ListConnectionDetails. */
public final class Main {
    /*
     * x-ms-original-file: specification/devspaces/resource-manager/Microsoft.DevSpaces/stable/2019-04-01/examples/ControllersListConnectionDetails_example.json
     */
    /**
     * Sample code: ControllersListConnectionDetails.
     *
     * @param manager Entry point to DevSpacesManager.
     */
    public static void controllersListConnectionDetails(com.azure.resourcemanager.devspaces.DevSpacesManager manager) {
        manager
            .controllers()
            .listConnectionDetailsWithResponse(
                "myResourceGroup",
                "myControllerResource",
                new ListConnectionDetailsParameters()
                    .withTargetContainerHostResourceId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerService/managedClusters/myCluster"),
                com.azure.core.util.Context.NONE);
    }
}
