package week02

/*不应该在Dao层Wrap错误
  原因: sql.ErrNoRows 表示没有在数据库中查到数据
  对调用业务方来说代表正常 or 非正常，需要业务方来做处理，
  应该返回空数组.
*/

