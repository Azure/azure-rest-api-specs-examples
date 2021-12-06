Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-desktopvirtualization_1.0.0-beta.1/sdk/desktopvirtualization/azure-resourcemanager-desktopvirtualization/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for UserSessions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/UserSession_Get.json
     */
    /**
     * Sample code: UserSession_Get.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void userSessionGet(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager
            .userSessions()
            .getWithResponse("resourceGroup1", "hostPool1", "sessionHost1.microsoft.com", "1", Context.NONE);
    }
}
```
