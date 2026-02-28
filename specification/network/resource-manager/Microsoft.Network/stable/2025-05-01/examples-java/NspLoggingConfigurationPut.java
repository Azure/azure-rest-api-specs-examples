
import com.azure.resourcemanager.network.fluent.models.NspLoggingConfigurationInner;
import java.util.Arrays;

/**
 * Samples for NetworkSecurityPerimeterLoggingConfigurations CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/NspLoggingConfigurationPut.
     * json
     */
    /**
     * Sample code: NspLoggingConfigurationPut.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void nspLoggingConfigurationPut(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkSecurityPerimeterLoggingConfigurations()
            .createOrUpdateWithResponse("rg1", "nsp1", "instance",
                new NspLoggingConfigurationInner().withEnabledLogCategories(
                    Arrays.asList("NspPublicInboundPerimeterRulesDenied", "NspPublicOutboundPerimeterRulesDenied")),
                com.azure.core.util.Context.NONE);
    }
}
