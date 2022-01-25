Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mysql_1.0.2/sdk/mysql/azure-resourcemanager-mysql/README.md) on how to add the SDK to your project and authenticate.

```java
/** Samples for Databases CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/DatabaseCreate.json
     */
    /**
     * Sample code: DatabaseCreate.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void databaseCreate(com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager
            .databases()
            .define("db1")
            .withExistingServer("TestGroup", "testserver")
            .withCharset("utf8")
            .withCollation("utf8_general_ci")
            .create();
    }
}
```
