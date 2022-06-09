```java
import com.azure.core.util.Context;

/** Samples for UserConfirmationPassword Send. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementUserConfirmationPasswordSend.json
     */
    /**
     * Sample code: ApiManagementUserConfirmationPasswordSend.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementUserConfirmationPasswordSend(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .userConfirmationPasswords()
            .sendWithResponse("rg1", "apimService1", "57127d485157a511ace86ae7", null, Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.
