
import com.azure.resourcemanager.network.fluent.models.CustomIpPrefixInner;

/**
 * Samples for CustomIpPrefixes CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/CustomIpPrefixCreateCustomizedValues.json
     */
    /**
     * Sample code: Create custom IP prefix allocation method.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void createCustomIPPrefixAllocationMethod(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getCustomIpPrefixes().createOrUpdate("rg1", "test-customipprefix",
            new CustomIpPrefixInner().withLocation("westus").withCidr("0.0.0.0/24"), com.azure.core.util.Context.NONE);
    }
}
