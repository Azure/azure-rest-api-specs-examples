
import com.azure.core.management.SubResource;
import com.azure.resourcemanager.network.fluent.models.NspAssociationInner;
import com.azure.resourcemanager.network.models.AssociationAccessMode;

/**
 * Samples for NetworkSecurityPerimeterAssociations CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/NspAssociationPut.json
     */
    /**
     * Sample code: NspAssociationPut.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void nspAssociationPut(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkSecurityPerimeterAssociations().createOrUpdate("rg1",
            "nsp1", "association1",
            new NspAssociationInner().withPrivateLinkResource(new SubResource().withId(
                "/subscriptions/{paasSubscriptionId}/resourceGroups/{paasResourceGroupName}/providers/{providerName}/{resourceType}/{resourceName}"))
                .withProfile(new SubResource().withId(
                    "/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityPerimeters/nsp1/profiles/{profileName}"))
                .withAccessMode(AssociationAccessMode.ENFORCED),
            com.azure.core.util.Context.NONE);
    }
}
