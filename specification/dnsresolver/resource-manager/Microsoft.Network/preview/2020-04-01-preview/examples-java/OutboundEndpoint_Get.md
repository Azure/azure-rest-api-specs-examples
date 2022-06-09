```java
import com.azure.core.util.Context;

/** Samples for OutboundEndpoints Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/dnsresolver/resource-manager/Microsoft.Network/preview/2020-04-01-preview/examples/OutboundEndpoint_Get.json
     */
    /**
     * Sample code: Retrieve outbound endpoint for DNS resolver.
     *
     * @param manager Entry point to DnsResolverManager.
     */
    public static void retrieveOutboundEndpointForDNSResolver(
        com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager
            .outboundEndpoints()
            .getWithResponse("sampleResourceGroup", "sampleDnsResolver", "sampleOutboundEndpoint", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-dnsresolver_1.0.0-beta.1/sdk/dnsresolver/azure-resourcemanager-dnsresolver/README.md) on how to add the SDK to your project and authenticate.
