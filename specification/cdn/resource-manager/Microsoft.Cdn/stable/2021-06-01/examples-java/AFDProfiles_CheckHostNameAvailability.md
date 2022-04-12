Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.14.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.cdn.models.CheckHostnameAvailabilityInput;

/** Samples for AfdProfiles CheckHostnameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/AFDProfiles_CheckHostNameAvailability.json
     */
    /**
     * Sample code: AFDProfiles_CheckHostNameAvailability.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void aFDProfilesCheckHostNameAvailability(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cdnProfiles()
            .manager()
            .serviceClient()
            .getAfdProfiles()
            .checkHostnameAvailabilityWithResponse(
                "RG",
                "profile1",
                new CheckHostnameAvailabilityInput().withHostname("www.someDomain.net"),
                Context.NONE);
    }
}
```
