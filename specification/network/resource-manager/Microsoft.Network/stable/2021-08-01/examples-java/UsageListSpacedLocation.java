import com.azure.core.util.Context;

/** Samples for Usages List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/UsageListSpacedLocation.json
     */
    /**
     * Sample code: List usages spaced location.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listUsagesSpacedLocation(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getUsages().list("West US", Context.NONE);
    }
}
