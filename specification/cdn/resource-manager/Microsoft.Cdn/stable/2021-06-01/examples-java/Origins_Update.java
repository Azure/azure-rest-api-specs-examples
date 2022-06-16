import com.azure.core.util.Context;
import com.azure.resourcemanager.cdn.models.OriginUpdateParameters;

/** Samples for Origins Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Origins_Update.json
     */
    /**
     * Sample code: Origins_Update.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void originsUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cdnProfiles()
            .manager()
            .serviceClient()
            .getOrigins()
            .update(
                "RG",
                "profile1",
                "endpoint1",
                "www-someDomain-net",
                new OriginUpdateParameters()
                    .withHttpPort(42)
                    .withHttpsPort(43)
                    .withOriginHostHeader("www.someDomain2.net")
                    .withPriority(1)
                    .withWeight(50)
                    .withEnabled(true)
                    .withPrivateLinkAlias(
                        "APPSERVER.d84e61f0-0870-4d24-9746-7438fa0019d1.westus2.azure.privatelinkservice"),
                Context.NONE);
    }
}
