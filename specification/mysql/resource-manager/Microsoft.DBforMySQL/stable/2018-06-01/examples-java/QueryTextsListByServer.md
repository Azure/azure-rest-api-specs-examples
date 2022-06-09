```java
import com.azure.core.util.Context;
import java.util.Arrays;

/** Samples for QueryTexts ListByServer. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2018-06-01/examples/QueryTextsListByServer.json
     */
    /**
     * Sample code: QueryTextsListByServer.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void queryTextsListByServer(com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager
            .queryTexts()
            .listByServer("testResourceGroupName", "testServerName", Arrays.asList("1", "2"), Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mysql_1.0.2/sdk/mysql/azure-resourcemanager-mysql/README.md) on how to add the SDK to your project and authenticate.
