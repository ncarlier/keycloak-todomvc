/*
 * Copyright 2016 (C) Worldline - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 */
package com.worldline.fe.todo;

import org.springframework.data.repository.PagingAndSortingRepository;
import org.springframework.data.rest.core.annotation.RepositoryRestResource;

/**
 * Todo REST repository.
 * @author Nicolas Carlier <nicolas.carlier@worldline.com>
 */
@RepositoryRestResource(collectionResourceRel = "todos", path = "todos")
public interface TodoRestRepository extends PagingAndSortingRepository<TodoModel, Long> {
  
}
