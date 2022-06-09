```java
import com.azure.core.util.Context;

/** Samples for UserSessions Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/UserSession_Delete.json
     */
    /**
     * Sample code: UserSession_Delete.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void userSessionDelete(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager
            .userSessions()
            .deleteWithResponse("resourceGroup1", "hostPool1", "sessionHost1.microsoft.com", "1", true, Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-desktopvirtualization_1.0.0-beta.1/sdk/desktopvirtualization/azure-resourcemanager-desktopvirtualization/README.md) on how to add the SDK to your project and authenticate.
