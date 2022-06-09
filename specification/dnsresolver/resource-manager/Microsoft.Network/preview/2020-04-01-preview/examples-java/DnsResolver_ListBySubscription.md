```java
import com.azure.core.util.Context;

/** Samples for DnsResolvers List. */
public final class Main {
    /*
     * x-ms-original-file: specification/dnsresolver/resource-manager/Microsoft.Network/preview/2020-04-01-preview/examples/DnsResolver_ListBySubscription.json
     */
    /**
     * Sample code: List DNS resolvers by subscription.
     *
     * @param manager Entry point to DnsResolverManager.
     */
    public static void listDNSResolversBySubscription(
        com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager.dnsResolvers().list(null, Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-dnsresolver_1.0.0-beta.1/sdk/dnsresolver/azure-resourcemanager-dnsresolver/README.md) on how to add the SDK to your project and authenticate.
