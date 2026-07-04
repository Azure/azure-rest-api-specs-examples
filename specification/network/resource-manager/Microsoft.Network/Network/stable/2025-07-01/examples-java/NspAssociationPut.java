
import com.azure.core.management.SubResource;
import com.azure.resourcemanager.network.fluent.models.NspAssociationInner;
import com.azure.resourcemanager.network.models.AssociationAccessMode;

/**
 * Samples for NetworkSecurityPerimeterAssociations CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NspAssociationPut.json
     */
    /**
     * Sample code: NspAssociationPut.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void nspAssociationPut(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkSecurityPerimeterAssociations().createOrUpdate("rg1", "nsp1", "association1",
            new NspAssociationInner().withPrivateLinkResource(new SubResource().withId(
                "/subscriptions/{paasSubscriptionId}/resourceGroups/{paasResourceGroupName}/providers/{providerName}/{resourceType}/{resourceName}"))
                .withProfile(new SubResource().withId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityPerimeters/nsp1/profiles/{profileName}"))
                .withAccessMode(AssociationAccessMode.ENFORCED),
            com.azure.core.util.Context.NONE);
    }
}
