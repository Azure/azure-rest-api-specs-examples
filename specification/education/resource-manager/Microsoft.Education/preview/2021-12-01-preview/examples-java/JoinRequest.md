Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-education_1.0.0-beta.1/sdk/education/azure-resourcemanager-education/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for JoinRequests Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/education/resource-manager/Microsoft.Education/preview/2021-12-01-preview/examples/JoinRequest.json
     */
    /**
     * Sample code: JoinRequest.
     *
     * @param manager Entry point to EducationManager.
     */
    public static void joinRequest(com.azure.resourcemanager.education.EducationManager manager) {
        manager
            .joinRequests()
            .getWithResponse(
                "{billingAccountName}",
                "{billingProfileName}",
                "{invoiceSectionName}",
                "{joinRequestName}",
                Context.NONE);
    }
}
```
