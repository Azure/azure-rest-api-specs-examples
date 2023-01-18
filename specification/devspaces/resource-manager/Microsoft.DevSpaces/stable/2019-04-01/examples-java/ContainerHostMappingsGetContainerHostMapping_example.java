import com.azure.resourcemanager.devspaces.fluent.models.ContainerHostMappingInner;

/** Samples for ContainerHostMappings GetContainerHostMapping. */
public final class Main {
    /*
     * x-ms-original-file: specification/devspaces/resource-manager/Microsoft.DevSpaces/stable/2019-04-01/examples/ContainerHostMappingsGetContainerHostMapping_example.json
     */
    /**
     * Sample code: ContainerHostMappingsGetContainerHostMapping.
     *
     * @param manager Entry point to DevSpacesManager.
     */
    public static void containerHostMappingsGetContainerHostMapping(
        com.azure.resourcemanager.devspaces.DevSpacesManager manager) {
        manager
            .containerHostMappings()
            .getContainerHostMappingWithResponse(
                "myResourceGroup",
                "eastus",
                new ContainerHostMappingInner()
                    .withContainerHostResourceId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerService/managedClusters/myCluster"),
                com.azure.core.util.Context.NONE);
    }
}
