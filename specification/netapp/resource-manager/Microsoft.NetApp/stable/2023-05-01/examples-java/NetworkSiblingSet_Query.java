import com.azure.resourcemanager.netapp.models.QueryNetworkSiblingSetRequest;

/** Samples for NetAppResource QueryNetworkSiblingSet. */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2023-05-01/examples/NetworkSiblingSet_Query.json
     */
    /**
     * Sample code: NetworkSiblingSet_Query.
     *
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void networkSiblingSetQuery(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager
            .netAppResources()
            .queryNetworkSiblingSetWithResponse(
                "eastus",
                new QueryNetworkSiblingSetRequest()
                    .withNetworkSiblingSetId("9760acf5-4638-11e7-9bdb-020073ca3333")
                    .withSubnetId(
                        "/subscriptions/9760acf5-4638-11e7-9bdb-020073ca7778/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testVnet/subnets/testSubnet"),
                com.azure.core.util.Context.NONE);
    }
}
